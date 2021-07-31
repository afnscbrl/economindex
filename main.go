package main

import (
	"economindex/scraps"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseGlob("index.html"))
	data := scraps.Scraping()
	tmpl.ExecuteTemplate(w, "Index", data)
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)
}
