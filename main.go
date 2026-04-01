package main

import (
	"fmt"
	"net/http"
	"os"

	"ascii/utils/api" // Fixed: correct module path
)

func main() {
	// Static files - fixed path
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("utils/api/templates/static"))))

	// Routes
	http.HandleFunc("/", api.HomeHandler)
	http.HandleFunc("/generate-qr", api.GenerateQRHandler)
	http.HandleFunc("/generate-logo", api.GenerateLogoHandler)
	http.HandleFunc("/download/ascii", api.DownloadASCII)
	http.HandleFunc("/download/qr-image", api.DownloadQRImage)
	http.HandleFunc("/api/qr", api.APIQR)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on http://localhost:%s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
