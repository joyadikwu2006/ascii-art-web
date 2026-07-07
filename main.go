package main

import (
	"ascii-art-web/asciiart"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 - Page Not Found", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, nil)
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	text := r.FormValue("text")
	bannerStyle := r.FormValue("bannerStyle")

	if text == "" {
		http.Error(w, "400 - Text cannot be empty", http.StatusBadRequest)
		return
	}
	bannerFile := "banners/" + bannerStyle + ".txt"
	banner, err := asciiart.LoadBanner(bannerFile)
	if err != nil {
		http.Error(w, "404 - Banner not found", http.StatusNotFound)
		return
	}
	result := asciiart.GenerateArt(text, banner)
	err = tmpl.Execute(w, map[string]string{
		"AsciiArt":    result,
		"Text":        text,
		"BannerStyle": bannerStyle,
	})
	if err != nil {
		http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
		return
	}
}
