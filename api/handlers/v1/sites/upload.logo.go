package sites

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"io"
	"net/http"
	"strconv"

	"golang.org/x/image/webp"

	"github.com/disintegration/imaging"
	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/actions"
	"github.com/ralphferrara/aria/base/agnostic"
	"github.com/ralphferrara/aria/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| UploadHandler :: Uploads or clears logo, stores in MinIO, updates site_logo
//||------------------------------------------------------------------------------------------------||

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	//||------------------------------------------------------------------------------------------------||
	//|| Get Session
	//||------------------------------------------------------------------------------------------------||

	_, account, _, err := actions.LoadSessionAccount(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Parse site ID
	//||------------------------------------------------------------------------------------------------||
	siteIDStr := r.FormValue("siteId")
	if siteIDStr == "" {
		responses.Error(w, http.StatusBadRequest, "Missing siteId")
		return
	}

	siteID, err := strconv.Atoi(siteIDStr)
	if err != nil || siteID <= 0 {
		responses.Error(w, http.StatusBadRequest, "Invalid siteId")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verify ownership
	//||------------------------------------------------------------------------------------------------||

	var count int64
	err = app.SQLDB["main"].DB.
		Table("sites").
		Where("id_site = ? AND fid_account = ? AND site_status NOT IN ?", siteID, account.ID, []string{"RMVD", "BNND"}).
		Count(&count).Error
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Database error")
		return
	}
	if count == 0 {
		responses.Error(w, http.StatusForbidden, "Unauthorized or invalid site")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Handle Delete (no file uploaded)
	//||------------------------------------------------------------------------------------------------||

	file, _, err := r.FormFile("image")
	if err == http.ErrMissingFile {
		// clear DB field
		if err := app.SQLDB["main"].DB.
			Table("sites").
			Where("id_site = ? AND fid_account = ?", siteID, account.ID).
			Update("site_logo", "").Error; err != nil {
			responses.Error(w, http.StatusInternalServerError, "Failed to clear site_logo: "+err.Error())
			return
		}

		// optional: remove from MinIO (if old filename known, you could query first)
		// _ = app.Storages["sites"].Delete(oldFilename)

		responses.Success(w, http.StatusOK, map[string]any{
			"object": "",
		})
		return
	}
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Error reading image: "+err.Error())
		return
	}
	defer file.Close()

	//||------------------------------------------------------------------------------------------------||
	//|| Read Image
	//||------------------------------------------------------------------------------------------------||

	buf, err := io.ReadAll(file)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Failed to read image data: "+err.Error())
		return
	}

	// reject animated GIFs
	if gifImg, err := gif.DecodeAll(bytes.NewReader(buf)); err == nil && len(gifImg.Image) > 1 {
		responses.Error(w, http.StatusBadRequest, "Animated GIFs are not supported")
		return
	}

	// decode with fallback for WebP
	var srcImg image.Image
	srcImg, err = imaging.Decode(bytes.NewReader(buf), imaging.AutoOrientation(true))
	if err != nil {
		if webpImg, webpErr := webp.Decode(bytes.NewReader(buf)); webpErr == nil {
			srcImg = webpImg
		} else {
			responses.Error(w, http.StatusBadRequest, "Invalid image: "+err.Error())
			return
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Resize + Canvas
	//||------------------------------------------------------------------------------------------------||

	resizedImg := imaging.Resize(srcImg, 500, 0, imaging.Lanczos)
	canvas := image.NewRGBA(image.Rect(0, 0, 500, 500))
	draw.Draw(canvas, canvas.Bounds(), image.Transparent, image.Point{}, draw.Src)
	offset := image.Pt((500-resizedImg.Bounds().Dx())/2, (500-resizedImg.Bounds().Dy())/2)
	draw.Draw(canvas, resizedImg.Bounds().Add(offset), resizedImg, image.Point{}, draw.Over)

	// encode to webp
	webpBuf, err := agnostic.EncodeWebP(canvas)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "WebP encoding failed: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Upload to Storage
	//||------------------------------------------------------------------------------------------------||

	filename := LogoHash(siteID, int(account.ID))
	fmt.Println("Uploading logo to storage: " + filename)
	if err := app.Storages["sites"].Put(filename, webpBuf.Bytes()); err != nil {
		responses.Error(w, http.StatusInternalServerError, "Storage upload failed: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Update DB record
	//||------------------------------------------------------------------------------------------------||

	if err := app.SQLDB["main"].DB.
		Table("sites").
		Where("id_site = ? AND fid_account = ?", siteID, account.ID).
		Update("site_logo", filename).Error; err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to update site_logo: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Success
	//||------------------------------------------------------------------------------------------------||
	responses.Success(w, http.StatusOK, map[string]any{
		"object": filename,
	})
}
