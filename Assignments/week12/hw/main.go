package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

// User is exported so that it can be an embedded type in pageData
type User struct {
	First    string
	Email    string
	Password string
	LoggedIn bool
}

type pageData struct {
	User
	Title   string
	Heading string
}

var db = map[string]User{}

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/signupprocess", processSignUp)
	http.HandleFunc("/login", login)
	http.HandleFunc("/loginprocess", processLogin)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("user-session")
	if err != nil {
		c = &http.Cookie{}
	}

	u := User{}
	ok := true
	if u, ok = db[c.Value]; ok {
		u.LoggedIn = true
	}

	pd := pageData{
		User:    u,
		Title:   "Ukraine",
		Heading: "Welcome To Ukraine",
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", pd)
	if err != nil {
		log.Println("index.gohtml couldn't ExecuteTemplate", err)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("user-session")
	if err != nil {
		c = &http.Cookie{}
	}

	u := User{}
	ok := false
	if u, ok = db[c.Value]; ok {
		u.LoggedIn = true
	}

	pd := pageData{
		User:    u,
		Title:   "About Ukraine",
		Heading: "All About Ukraine",
	}

	err = tpl.ExecuteTemplate(w, "about.gohtml", pd)
	if err != nil {
		log.Println("about.gohtml couldn't ExecuteTemplate", err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("user-session")
	if err != nil {
		c = &http.Cookie{}
	}

	u := User{}
	ok := false
	if u, ok = db[c.Value]; ok {
		u.LoggedIn = true
	}

	pd := pageData{
		User:    u,
		Title:   "Cities",
		Heading: "Cities to Visit",
	}

	err = tpl.ExecuteTemplate(w, "contact.gohtml", pd)
	if err != nil {
		log.Println("contact.gohtml couldn't ExecuteTemplate", err)
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("user-session")
	if err != nil {
		c = &http.Cookie{}
	}

	u := User{}
	if _, ok := db[c.Value]; ok {
		// already logged in, so redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	pd := pageData{
		User:    u,
		Title:   "SIGN UP",
		Heading: "SIGN UP ",
	}

	err = tpl.ExecuteTemplate(w, "signup.gohtml", pd)
	if err != nil {
		log.Println("signup.gohtml couldn't ExecuteTemplate", err)
	}
}

func processSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusNotFound)
		return
	}

	f := r.FormValue("fn")
	if f == "" {
		http.Error(w, "first name cannot be empty", http.StatusBadRequest)
	}

	e := r.FormValue("em")
	if e == "" {
		http.Error(w, "email cannot be empty", http.StatusBadRequest)
	}

	p := r.FormValue("pw")
	if p == "" {
		http.Error(w, "password cannot be empty", http.StatusBadRequest)
	}

	u := User{
		First:    f,
		Email:    e,
		Password: p,
	}

	db[e] = u

	c := &http.Cookie{
		Name:  "user-session",
		Value: e,
		Path:  "/",
	}

	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func login(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("user-session")
	if err != nil {
		c = &http.Cookie{}
	}

	u := User{}
	if _, ok := db[c.Value]; ok {
		// already logged in, so redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	pd := pageData{
		User:    u,
		Title:   "LOGIN",
		Heading: "LOGIN",
	}

	tpl.ExecuteTemplate(w, "login.gohtml", pd)
}

func processLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusNotFound)
		return
	}

	e := r.FormValue("em")
	if e == "" {
		http.Error(w, "email cannot be empty", http.StatusBadRequest)
	}

	p := r.FormValue("pw")
	if p == "" {
		http.Error(w, "password cannot be empty", http.StatusBadRequest)
	}

	u := User{}
	ok := false
	if u, ok = db[e]; !ok {
		http.Error(w, "username or password not correct", http.StatusBadRequest)
		return
	}

	if p != u.Password {
		http.Error(w, "username or password not correct", http.StatusBadRequest)
		return
	}

	c := &http.Cookie{
		Name:  "user-session",
		Value: e,
		Path:  "/",
	}

	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("user-session")
	if err != nil {
		c = &http.Cookie{
			Name: "user-session",
		}
	}

	c.MaxAge = -1

	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
