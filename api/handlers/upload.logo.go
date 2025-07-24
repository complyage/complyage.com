package handlers

//||------------------------------------------------------------------------------------------------||
//|| Set Policy
//||------------------------------------------------------------------------------------------------||
// mc.exe anonymous set download babl/complyage.com/sites/logos

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db"
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"base/agnostic"
	"base/helpers"
	"base/responses"

	"github.com/disintegration/imaging"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

//||------------------------------------------------------------------------------------------------||
//|| UploadHandler :: Handles image upload and conversion to WebP
//||------------------------------------------------------------------------------------------------||

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Check Method
	//||------------------------------------------------------------------------------------------------||

	if r.Method != http.MethodPost {
		responses.Error(w, http.StatusMethodNotAllowed, "Only POST allowed")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Session Cookie
	//||------------------------------------------------------------------------------------------------||

	cookie, err := r.Cookie("session")
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "No session cookie")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Session Record
	//||------------------------------------------------------------------------------------------------||

	session, err := helpers.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Read Image into Memory
	//||------------------------------------------------------------------------------------------------||

	file, _, err := r.FormFile("image")
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Error reading image: "+err.Error())
		return
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Failed to read image data: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Fail on Animated GIFs
	//||------------------------------------------------------------------------------------------------||

	if gifImg, err := gif.DecodeAll(bytes.NewReader(buf)); err == nil && len(gifImg.Image) > 1 {
		responses.Error(w, http.StatusBadRequest, "Animated GIFs are not supported. Please upload a static image.")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Decode Image
	//||------------------------------------------------------------------------------------------------||

	srcImg, err := imaging.Decode(bytes.NewReader(buf), imaging.AutoOrientation(true))
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid image: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate logoHash (hash = siteID_hashValue)
	//||------------------------------------------------------------------------------------------------||

	logoHash := r.FormValue("logoHash")
	if logoHash == "" || !strings.Contains(logoHash, "_") {
		responses.Error(w, http.StatusBadRequest, "Missing or invalid logoHash")
		return
	}

	parts := strings.SplitN(logoHash, "_", 2)
	siteIDStr := parts[0]
	clientHash := parts[1]

	siteID, err := strconv.Atoi(siteIDStr)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid site ID in hash")
		return
	}

	secret := os.Getenv("MINIO_HASH")
	hashInput := fmt.Sprintf("%s%d%s", secret, siteID, secret)
	expectedHashBytes := sha256.Sum256([]byte(hashInput))
	expectedHash := hex.EncodeToString(expectedHashBytes[:])

	if clientHash != expectedHash {
		responses.Error(w, http.StatusUnauthorized, "logoHash verification failed")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verify Site Ownership and Status
	//||------------------------------------------------------------------------------------------------||

	var count int64
	err = db.DB.
		Table("sites").
		Where("id_site = ? AND fid_account = ? AND site_status NOT IN ?", siteID, session.ID, []string{"RMVD", "BNND"}).
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
	//|| Resize + Canvas
	//||------------------------------------------------------------------------------------------------||

	resizedImg := imaging.Resize(srcImg, 500, 0, imaging.Lanczos)

	canvas := image.NewRGBA(image.Rect(0, 0, 500, 500))
	draw.Draw(canvas, canvas.Bounds(), image.Transparent, image.Point{}, draw.Src)

	offset := image.Pt(
		(500-resizedImg.Bounds().Dx())/2,
		(500-resizedImg.Bounds().Dy())/2,
	)
	draw.Draw(canvas, resizedImg.Bounds().Add(offset), resizedImg, image.Point{}, draw.Over)

	//||------------------------------------------------------------------------------------------------||
	//|| Convert to WebP (via agnostic)
	//||------------------------------------------------------------------------------------------------||

	webpBuf, err := agnostic.EncodeWebP(canvas)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "WebP encoding failed: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| MinIO Configuration
	//||------------------------------------------------------------------------------------------------||

	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	useSSL := os.Getenv("MINIO_USE_SSL") == "true"
	bucketName := os.Getenv("VITE_MINIO_BUCKET")

	//||------------------------------------------------------------------------------------------------||
	//|| Upload to MinIO
	//||------------------------------------------------------------------------------------------------||

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "MinIO connection failed: "+err.Error())
		return
	}

	objectName := fmt.Sprintf("sites/logos/%s.webp", logoHash)

	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Bucket check failed: "+err.Error())
		return
	}

	if !exists {
		responses.Error(w, http.StatusBadRequest, "Bucket does not exist")
		return
	}

	_, err = minioClient.PutObject(ctx, bucketName, objectName, webpBuf, int64(webpBuf.Len()), minio.PutObjectOptions{
		ContentType: "image/webp",
	})
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Upload failed: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Success Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]any{
		"object": objectName,
	})
}
