package main

import (
	"embed"
	"fmt"
	"net/http"
	"os"

	"ascii/utils/api"
)

// Embed static files
//go:embed utils/api/templates/static/*
var staticFiles embed.FS

func main() {
	// Serve embedded static files
	fs := http.FS(staticFiles)
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(fs)))

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

	fmt.Printf("Server running on port %s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
