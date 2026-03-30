package main

import (
	"fmt"
	"net/http"

	"github.com/HACKWITHNESBITT/Ascii/utils/api"
)

func main() {
	// Static files
	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	// Routes
	http.HandleFunc("/", api.HomeHandler)
	http.HandleFunc("/generate", api.GenerateHandler)
	http.HandleFunc("/download/png", api.DownloadPNG)
	http.HandleFunc("/download/ascii", api.DownloadASCII)
	http.HandleFunc("/api/qr", api.APIQR)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
