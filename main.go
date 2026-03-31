package main

import (
	"fmt"
	"net/http"

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
	http.HandleFunc("/api/qr", api.APIQR)

	fmt.Println("Server running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
