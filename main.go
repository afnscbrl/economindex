package main

import (
	"economindex/scraps"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func redirect(w http.ResponseWriter, req *http.Request) {
	_host := strings.Split(req.Host, ":")
	_host[1] = "8443"

	target := "https://" + strings.Join(_host, ":") + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}

	http.Redirect(w, req, target, http.StatusTemporaryRedirect)
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
	go http.ListenAndServe(":"+port, http.HandlerFunc(redirect))
}
