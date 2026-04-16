package handler

import (
	"html/template"
	"log"
	"net/http"

	"ascii/utils"
)

var templates *template.Template

func init() {
	var err error
	templates, err = template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("failed to load template: %v", err)
	}
}

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
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Println("template error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
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

	err := templates.ExecuteTemplate(w, "index.html", struct {
		ASCIIQR string
		Logo    string
		Text    string
	}{
		ASCIIQR: asciiQR,
		Logo:    "",
		Text:    text,
	})

	if err != nil {
		log.Println("template error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
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

	err := templates.ExecuteTemplate(w, "index.html", struct {
		ASCIIQR string
		Logo    string
		Text    string
	}{
		ASCIIQR: "",
		Logo:    logo,
		Text:    text,
	})

	if err != nil {
		log.Println("template error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
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
