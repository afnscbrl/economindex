package main

import (
	"economindex/scraps"
	"html/template"
	"log"
	"net/http"
	"os"
)

// func redirect(w http.ResponseWriter, r *http.Request) {
// 	http.Redirect(w, r, "https://"+r.Host, http.StatusMovedPermanently)
// 	fmt.Println("host:", r.Host)
// }

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
	http.ListenAndServe(":"+port, nil)
}
