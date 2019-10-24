package main

import (
	"html/template"
	"net/http"
	// "text/template"
)

var tpl *template.Template

type pageData struct {
	Title    string
	Heading  string
	Students string
	Teacher  string
	HW1      string
	HW2      string
	HW3      string
}

func init() {
	tpl = template.Must(template.ParseGlob("./template/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/class", class)
	http.HandleFunc("/students", students)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	pd := pageData{
		Title:   "HOME",
		Heading: "Welcome to our site CIT90 Fall2019",
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", pd)
	if err != nil {
		http.Error(w, "couldn't execute template", http.StatusInternalServerError)
		return
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	pd := pageData{
		Title:   "ABOUT",
		Heading: "Home Works Assignments",
		HW1:     "1. Lorem ipsum dolor sit amet consectetur adipisicing elit. Delectus ex quasi magnam illum deserunt dolor rem aspernatur, quam, odit perspiciatis labore? Ut et temporibus deserunt, voluptate enim modi officiis totam?",
		HW2:     "2. Lorem ipsum dolor sit amet consectetur adipisicing elit. Delectus ex quasi magnam illum deserunt dolor rem aspernatur, quam, odit perspiciatis labore? Ut et temporibus deserunt, voluptate enim modi officiis totam?",
		HW3:     "3. Lorem ipsum dolor sit amet consectetur adipisicing elit. Delectus ex quasi magnam illum deserunt dolor rem aspernatur, quam, odit perspiciatis labore? Ut et temporibus deserunt, voluptate enim modi officiis totam?",
	}

	err := tpl.ExecuteTemplate(w, "about.gohtml", pd)
	if err != nil {
		http.Error(w, "couldn't execute template", http.StatusInternalServerError)
		return
	}
}

func class(w http.ResponseWriter, r *http.Request) {
	pd := pageData{
		Title:   "Class",
		Heading: "Reading, Writing",
		Teacher: "Mr. Richard Santod",
	}

	err := tpl.ExecuteTemplate(w, "class.gohtml", pd)
	if err != nil {
		http.Error(w, "couldn't execute template", http.StatusInternalServerError)
		return
	}

}
func students(w http.ResponseWriter, r *http.Request) {
	pd := pageData{
		Title:    "Students",
		Heading:  "Your list of classmates:",
		Students: "Rob, Lily, Kate, Sophy, Cindy, Mark, Nala",
	}

	err := tpl.ExecuteTemplate(w, "students.gohtml", pd)
	if err != nil {
		http.Error(w, "couldn't execute template", http.StatusInternalServerError)
		return
	}
}
