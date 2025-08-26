package images

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"

	"github.com/disintegration/imaging"
)

//||------------------------------------------------------------------------------------------------||
//|| ResizeBase64Image: Resize base64-encoded image to maxDim x maxDim (preserving aspect ratio)
//||------------------------------------------------------------------------------------------------||

func ResizeBase64Image(b64 string, maxDim int) (string, error) {
	// Decode base64 to bytes
	imgBytes, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 image: %w", err)
	}
	// Decode image
	img, format, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return "", fmt.Errorf("failed to decode image: %w", err)
	}
	// Resize (preserve aspect ratio)
	resized := imaging.Fit(img, maxDim, maxDim, imaging.Lanczos)
	// Encode resized image to buffer
	buf := new(bytes.Buffer)
	switch format {
	case "jpeg":
		err = imaging.Encode(buf, resized, imaging.JPEG)
	case "png":
		err = imaging.Encode(buf, resized, imaging.PNG)
	default:
		return "", fmt.Errorf("unsupported image format: %s", format)
	}
	if err != nil {
		return "", fmt.Errorf("failed to encode resized image: %w", err)
	}
	// Re-encode to base64
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}
