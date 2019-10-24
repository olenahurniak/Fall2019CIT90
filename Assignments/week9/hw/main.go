package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/catalog", catalog)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	xs := []string{"muvies", "musics", "games", "books", "programs"}
	err := tpl.ExecuteTemplate(w, "index.gohtml", xs)
	if err != nil {
		http.Error(w, "could not run template", http.StatusInternalServerError)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	m := map[int]string{}
	m[1] = "Lily"
	m[2] = "Armine"
	err := tpl.ExecuteTemplate(w, "about.gohtml", m)
	if err != nil {
		http.Error(w, "could not run template", http.StatusInternalServerError)
	}
}

func catalog(w http.ResponseWriter, r *http.Request) {
	m := map[int]string{}
	m[1] = "Lily"
	m[2] = "Armine"
	err := tpl.ExecuteTemplate(w, "catalog.gohtml", m)
	if err != nil {
		http.Error(w, "could not run template", http.StatusInternalServerError)
	}
}
