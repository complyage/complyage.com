package adapters

import (
	"base/verify"
	"bytes"
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"time"

	"github.com/fogleman/gg"
	"github.com/skip2/go-qrcode"
)

func GeneratePostcardBack(addr verify.Address, uuid string, checkCode string) ([]byte, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Generate the Verify URL
	//||------------------------------------------------------------------------------------------------||

	verifyURL := fmt.Sprintf("%s%s%s", os.Getenv("VITE_COMPLYAGE_UI_URL"), os.Getenv("VERIFICATION_ADDRESS_URL"), checkCode)
	fmt.Println("Verify URL:", verifyURL)

	//||------------------------------------------------------------------------------------------------||
	//|| Var
	//||------------------------------------------------------------------------------------------------||

	addrFontSize := 14.0
	leftMargin := 35   // distance from left edge
	rightMargin := -30 // distance from right edge
	topMargin := 40    // minimum distance from top edge
	boldFontPath := "./.assets/fonts/DejaVuLGCSans-Bold.ttf"
	regularFontPath := "./.assets/fonts/DejaVuLGCSans.ttf"

	//||------------------------------------------------------------------------------------------------||
	//|| Template File
	//||------------------------------------------------------------------------------------------------||

	templateFile, err := os.Open("./.assets/postcard-back.png")
	if err != nil {
		return []byte(""), fmt.Errorf("open template: %w", err)
	}
	defer templateFile.Close()
	templateImg, err := png.Decode(templateFile)
	if err != nil {
		return []byte(""), fmt.Errorf("decode template: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| QR Code
	//||------------------------------------------------------------------------------------------------||

	qrSize := 90
	qrPng, err := qrcode.Encode(verifyURL, qrcode.Medium, qrSize)
	if err != nil {
		return []byte(""), fmt.Errorf("encode qr: %w", err)
	}
	qrImg, err := png.Decode(bytes.NewReader(qrPng))
	if err != nil {
		return []byte(""), fmt.Errorf("decode qr: %w", err)
	}

	W := templateImg.Bounds().Dx()
	H := templateImg.Bounds().Dy()
	dc := gg.NewContext(W, H)
	dc.DrawImage(templateImg, 0, 0)

	totalBlockHeight := qrSize + 28 // 28px for code + spacing

	blockY := (H - totalBlockHeight) / 2
	if blockY < topMargin {
		blockY = topMargin
	}

	qrX := leftMargin + qrSize/2 + 25
	qrY := blockY + 25

	dc.DrawImage(qrImg, qrX-qrSize/2, qrY)

	//||------------------------------------------------------------------------------------------------||
	//|| Code Block
	//||------------------------------------------------------------------------------------------------||

	dc.SetRGB(0, 0, 0)
	codeX := float64(qrX)
	codeY := float64(qrY + qrSize + 20)
	codeZ := float64(qrY + qrSize + 40)

	if err := dc.LoadFontFace(regularFontPath, 12); err != nil {
		return []byte(""), fmt.Errorf("font: %w", err)
	}

	dc.DrawStringAnchored("Your verification code is", codeX, codeY, 0.5, 0)

	if err := dc.LoadFontFace(boldFontPath, 18); err != nil {
		return []byte(""), fmt.Errorf("font: %w", err)
	}

	dc.DrawStringAnchored(checkCode, codeX, codeZ, 0.5, 0)

	//||------------------------------------------------------------------------------------------------||
	//|| Address Block
	//||------------------------------------------------------------------------------------------------||

	if "SET_ADDRESS" == "NOTHANKS" {
		lines := 4.0
		textHeight := addrFontSize * lines * 1.3
		addrX := float64(W - rightMargin)
		addrY := float64(H)/2 - textHeight/2 + 10

		if err := dc.LoadFontFace(boldFontPath, 10); err != nil {
			return []byte(""), fmt.Errorf("font: %w", err)
		}
		dc.SetRGB(0, 0, 0)
		dc.DrawStringAnchored("Deliver to :", (addrX - 140), (addrY - 14), 0.5, 0.5)

		if err := dc.LoadFontFace(regularFontPath, addrFontSize); err != nil {
			return []byte(""), fmt.Errorf("font: %w", err)
		}
		addressBlock := fmt.Sprintf("%s\n%s\n%s, %s %s\n%s", addr.Line1, addr.Line2, addr.City, addr.State, addr.Postal, addr.Country)
		dc.SetRGB(0, 0, 0)
		dc.DrawStringWrapped(addressBlock, addrX, addrY, 1.0, 0, 280, 1.3, gg.AlignCenter)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| UUID BLACK BOX AT BOTTOM
	//||------------------------------------------------------------------------------------------------||

	boxHeight := 20.0
	boxY := float64(H) - boxHeight - 5
	dc.SetRGB(0, 0, 0)
	dc.DrawRectangle(5, boxY, float64(W-10), boxHeight)
	dc.Fill()

	uuidFontSize := 10.0
	if err := dc.LoadFontFace(boldFontPath, uuidFontSize); err != nil {
		return []byte(""), fmt.Errorf("font: %w", err)
	}
	dc.SetRGB(1, 1, 1) // white
	dc.DrawStringAnchored(uuid, float64(W)/2, boxY+boxHeight/2, 0.5, 0.5)

	//||------------------------------------------------------------------------------------------------||
	//|| Save
	//||------------------------------------------------------------------------------------------------||

	if "SAVE_TO_DISK" == "NOTHANKS" {
		tmpDir := "./tmp"
		if err := os.MkdirAll(tmpDir, 0755); err != nil {
			return []byte(""), fmt.Errorf("create tmp dir: %w", err)
		}
		outPath := filepath.Join(tmpDir, fmt.Sprintf("postcard-back-%d.png", time.Now().UnixNano()))
		outFile, err := os.Create(outPath)
		if err != nil {
			return []byte(""), fmt.Errorf("create out: %w", err)
		}
		defer outFile.Close()
		if err := png.Encode(outFile, dc.Image()); err != nil {
			return []byte(""), fmt.Errorf("encode png: %w", err)
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return as a Buffer
	//||------------------------------------------------------------------------------------------------||

	var buf bytes.Buffer
	if err := png.Encode(&buf, dc.Image()); err != nil {
		return nil, fmt.Errorf("encode png: %w", err)
	}
	return buf.Bytes(), nil

}
