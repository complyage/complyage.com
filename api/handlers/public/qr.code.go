package public

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/skip2/go-qrcode"
)

// QRCodeHandler : HTTP handler that generates a QR code from a base64 string
func QRCodeHandler(w http.ResponseWriter, r *http.Request) {
	// Get ?data= param
	encoded := r.URL.Query().Get("data")
	if encoded == "" {
		http.Error(w, "missing data parameter", http.StatusBadRequest)
		return
	}

	// Decode from base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid base64: %v", err), http.StatusBadRequest)
		return
	}

	// Generate QR code PNG (256px size, medium error correction)
	png, err := qrcode.Encode(string(decoded), qrcode.Medium, 256)
	if err != nil {
		http.Error(w, fmt.Sprintf("qr encode error: %v", err), http.StatusInternalServerError)
		return
	}

	// Write PNG to response
	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(png)
}
