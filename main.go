package main

import (
	"economindex/scraps"
	"html/template"
	"log"
	"net/http"
	"os"
)

func redirect(w http.ResponseWriter, req *http.Request) {
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target,
		// see comments below and consider the codes 308, 302, or 301
		http.StatusTemporaryRedirect)
}

func Index(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseGlob("index.html"))
	data := scraps.Scraping()
	tmpl.ExecuteTemplate(w, "Index", data)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", Index)
	http.ListenAndServe(":"+port, http.HandlerFunc(redirect))
}
