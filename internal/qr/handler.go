package qr

import (
	"fmt"
	"log"
	"net/http"

	"github.com/skip2/go-qrcode"
)

// GenerateQRCode handles the HTTP request to generate a QR code
func GenerateQRCode(w http.ResponseWriter, r *http.Request) {
	// For simplicity, let's use hardcoded user details
	userDetails := "User Name: John Doe\nLicense Class: A\nLicense ID: 123456789"

	// Generate the QR code
	png, err := qrcode.Encode(userDetails, qrcode.Medium, 256)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error generating QR code: %v", err), http.StatusInternalServerError)
		log.Printf("Error generating QR code: %v", err)
		return
	}

	// Set the response header to PNG
	w.Header().Set("Content-Type", "image/png")
	w.Write(png) // Write the QR code PNG image to the response
}
