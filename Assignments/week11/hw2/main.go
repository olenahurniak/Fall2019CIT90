package main

import (
	"html/template"
	"net/http"
)

// PageData some data
type PageData struct {
	Title string
	Heading string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/shop", shop)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	d := PageData{
		Title: "Ukraine",
		Heading: "My Home",
	}
	tpl.ExecuteTemplate(w, "index.gohtml", d)
}
func about(w http.ResponseWriter, r *http.Request) {
	d := PageData{
		Title: "About",
		Heading: "About My Country",
	}
	tpl.ExecuteTemplate(w, "about.gohtml", d)

}
func shop(w http.ResponseWriter, r *http.Request) {
	d := PageData{
		Title: "PLACES",
		Heading: "Places to GO",
	}
	tpl.ExecuteTemplate(w, "shop.gohtml", d)
}
func contact(w http.ResponseWriter, r *http.Request) {
	d := PageData{
		Title: "Contact",
		Heading: "Contacts",
	}
	tpl.ExecuteTemplate(w, "contact.gohtml", d)
}
