package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/skip2/go-qrcode"
)

var templates = template.Must(template.ParseFiles("utils/api/templates/index.html"))

// GenerateASCII generates ASCII QR from text
func GenerateASCII(text string) string {
	qr, err := qrcode.New(text, qrcode.Medium)
	if err != nil {
		return "Error generating QR"
	}
	bitmap := qr.Bitmap()
	ascii := ""
	for _, row := range bitmap {
		for _, col := range row {
			if col {
				ascii += "██"
			} else {
				ascii += "  "
			}
		}
		ascii += "\n"
	}
	return ascii
}

// Handlers
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	text := r.FormValue("text")
	if text == "" {
		http.Error(w, "Text required", http.StatusBadRequest)
		return
	}
	ascii := GenerateASCII(text)
	templates.ExecuteTemplate(w, "index.html", struct {
		ASCII string
		Text  string
	}{
		ASCII: ascii,
		Text:  text,
	})
}

func DownloadPNG(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "Text required", http.StatusBadRequest)
		return
	}
	qr, err := qrcode.New(text, qrcode.Medium)
	if err != nil {
		http.Error(w, "QR generation failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Disposition", "attachment; filename=qr.png")
	qr.PNG(256)
}

func DownloadASCII(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "Text required", http.StatusBadRequest)
		return
	}
	ascii := GenerateASCII(text)
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", "attachment; filename=qr-ascii.txt")
	w.Write([]byte(ascii))
}

func APIQR(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "Text required", http.StatusBadRequest)
		return
	}
	ascii := GenerateASCII(text)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"ascii":"` + template.HTMLEscapeString(ascii) + `"}`))
}

func main() {
	// Static files - fixed path
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("utils/api/templates/static"))))
	// Routes
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/generate", GenerateHandler)
	http.HandleFunc("/download/png", DownloadPNG)
	http.HandleFunc("/download/ascii", DownloadASCII)
	http.HandleFunc("/api/qr", APIQR)

	fmt.Println("Server running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
