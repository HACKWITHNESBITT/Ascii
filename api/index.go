package handler

import (
	"html/template"
	"net/http"

	"ascii/utils"
)

var templates = template.Must(template.ParseFiles("utils/api/templates/index.html"))

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		HomeHandler(w, r)
	case "/generate-qr":
		GenerateQRHandler(w, r)
	case "/generate-logo":
		GenerateLogoHandler(w, r)
	case "/download/ascii":
		DownloadASCII(w, r)
	case "/download/qr-image":
		DownloadQRImage(w, r)
	case "/api/qr":
		APIQR(w, r)
	default:
		http.NotFound(w, r)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func GenerateQRHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	if text == "" {
		http.Error(w, "Text required", http.StatusBadRequest)
		return
	}

	asciiQR := utils.GenerateASCII(text)
	templates.ExecuteTemplate(w, "index.html", struct {
		ASCIIQR string
		Logo    string
		Text    string
	}{ASCIIQR: asciiQR, Logo: "", Text: text})
}

func GenerateLogoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	if text == "" {
		http.Error(w, "Text required", http.StatusBadRequest)
		return
	}

	logo := utils.GenerateLogo(text)
	templates.ExecuteTemplate(w, "index.html", struct {
		ASCIIQR string
		Logo    string
		Text    string
	}{ASCIIQR: "", Logo: logo, Text: text})
}

func DownloadASCII(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "Text required", http.StatusBadRequest)
		return
	}

	ascii := utils.GenerateASCII(text)
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", "attachment; filename=qr-ascii.txt")
	w.Write([]byte(ascii))
}

func DownloadQRImage(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "Text required", http.StatusBadRequest)
		return
	}

	pngData, err := utils.GeneratePNG(text)
	if err != nil {
		http.Error(w, "Error generating QR image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Disposition", "attachment; filename=qr-code.png")
	w.Write(pngData)
}

func APIQR(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "Text required", http.StatusBadRequest)
		return
	}

	ascii := utils.GenerateASCII(text)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"ascii":"` + template.HTMLEscapeString(ascii) + `"}`))
}
