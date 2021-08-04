package main

import (
	"economindex/scraps"
	"html/template"
	"log"
	"net/http"
	"os"
)

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	// remove/add not default ports from req.Host
	target := "https://" + r.Host + r.RequestURI
	if len(r.URL.RawQuery) > 0 {
		target += "?" + r.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, r, target,
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
	// redirect every http request to https
	http.ListenAndServe(":"+port, http.HandlerFunc(redirectToHttps))
	// serve index (and anything else) as https
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	http.ListenAndServeTLS(":"+port, "cert.pem", "key.pem", mux)
}
