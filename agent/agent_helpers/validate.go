package agent_helpers

import (
	"agent/agent_interfaces"
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	_ "image/gif" // optional
	_ "image/jpeg"
	_ "image/png"
	"net/url"
	"strings"

	_ "golang.org/x/image/bmp" // for BMP
	_ "golang.org/x/image/webp"
)

//||------------------------------------------------------------------------------------------------||
//|| Allowed Types
//||------------------------------------------------------------------------------------------------||

var AllowedMIMEs = map[string]bool{
	"image/jpeg":         true,
	"image/png":          true,
	"image/webp":         true,
	"image/bmp":          true,
	"image/vnd.wap.wbmp": true, // WBMP MIME type
}

//||------------------------------------------------------------------------------------------------||
//|| ValidateMedia :: Checks MIME type and Base64 validity
//||------------------------------------------------------------------------------------------------||

func ValidateMedia(m agent_interfaces.AgentMedia) error {
	// 1. Check MIME type
	if _, ok := AllowedMIMEs[m.Mime]; !ok {
		return fmt.Errorf("unsupported MIME type: %s", m.Mime)
	}

	// 2. Ensure at least one field is provided
	if m.Blob == "" && m.Base64 == "" {
		return errors.New("media must contain either a blob or base64")
	}

	// 3. If Base64 is provided, validate it
	if m.Base64 != "" {
		data, err := base64.StdEncoding.DecodeString(m.Base64)
		if err != nil {
			return errors.New("invalid base64 encoding")
		}

		// 4. Validate image format for everything except WBMP
		if m.Mime != "image/vnd.wap.wbmp" {
			if _, _, err := image.Decode(bytes.NewReader(data)); err != nil {
				return errors.New("decoded data is not a valid image")
			}
		}
	}

	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| ValidateCallbackURL
//||------------------------------------------------------------------------------------------------||

func ValidateCallbackURL(cb string) error {
	if cb == "" {
		return errors.New("callback URL is required")
	}

	// Parse URL
	parsed, err := url.Parse(cb)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return errors.New("invalid callback URL")
	}

	// Only allow http(s)
	scheme := strings.ToLower(parsed.Scheme)
	if scheme != "http" && scheme != "https" {
		return errors.New("callback URL must use http or https")
	}

	// (Optional) Restrict to certain domains
	// allowedDomain := "example.com"
	// if !strings.HasSuffix(parsed.Hostname(), allowedDomain) {
	//     return errors.New("callback URL not allowed")
	// }

	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| Requires Image
//||------------------------------------------------------------------------------------------------||

func RequiresImage(model string, action string) int {
	marker := strings.ToLower(model) + "." + strings.ToLower(action)
	switch marker {
	case "face.compare":
		return 2
	case "face.detect":
		return 1
	case "nsfw.detect":
		return 1
	case "ocr.extract":
		return 1
	case "vision.analyze":
		return 1
	case "vision.content":
		return 0
	default:
		return 0
	}
}

//||------------------------------------------------------------------------------------------------||
//|| Check Valid Model
//||------------------------------------------------------------------------------------------------||

func IsValidModel(model string) bool {
	for _, m := range agent_interfaces.AllowAgentModels {
		if m == model {
			return true
		}
	}
	return false
}

//||------------------------------------------------------------------------------------------------||
//|| Check if a Valid Action
//||------------------------------------------------------------------------------------------------||

func IsValidAction(model string, action string) bool {
	marker := strings.ToLower(model) + "." + strings.ToLower(action)
	for _, a := range agent_interfaces.AllowedAgentActions {
		if a == marker {
			return true
		}
	}
	return false
}
