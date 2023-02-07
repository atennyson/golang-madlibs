package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))

	})

	if name := r.FormValue("name"); name != "" {
		log.Printf("Name: %s", r.FormValue("name"))
	}
	if compliment := r.FormValue("compliment"); compliment != "" {
		log.Printf("Compliment: %s", r.FormValue("compliment"))
	}

	t.templ.Execute(w, r)
}

func main() {
	http.Handle("/", &templateHandler{filename: "index.html"})
	http.Handle("/greet", &templateHandler{filename: "greeting.html"})
	http.Handle("/words", &templateHandler{filename: "words.html"})
	http.Handle("/madlib", &templateHandler{filename: "madlib.html"})
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Ran into an error:", err)
	}
}
