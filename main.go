package main

import (
	"net/http"

	"ascii/utils/api"
)

// Vercel entrypoint handler
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		api.HomeHandler(w, r)

	case "/generate-qr":
		api.GenerateQRHandler(w, r)

	case "/generate-logo":
		api.GenerateLogoHandler(w, r)

	case "/download/ascii":
		api.DownloadASCII(w, r)

	case "/download/qr-image":
		api.DownloadQRImage(w, r)

	case "/api/qr":
		api.APIQR(w, r)

	default:
		http.NotFound(w, r)
	}
}
