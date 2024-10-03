package main

import (
	// Use the correct module name
	"fmt"
	"log"
	"net/http"

	"github.com/Heshan-Lahiru/QR-police-backend/internal/qr"
)

func main() {
	fmt.Println("Starting QR Code Generator Backend...")
	http.HandleFunc("/generate-qr", qr.GenerateQRCode) // Route to handle QR code generation
	log.Fatal(http.ListenAndServe(":8080", nil))       // Start the server on port 8080
}
