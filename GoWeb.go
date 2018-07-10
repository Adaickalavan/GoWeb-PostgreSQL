package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"fmt"
)

func viewhandler(w http.ResponseWriter, r *http.Request, title string) {
	title, err := getTitle(w, r)
	if err != nil {
		fmt.Println("inside viewhandler",err)
		return
	}
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func edithandler(w http.ResponseWriter, r *http.Request, title string) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func savehandler(w http.ResponseWriter, r *http.Request, title string) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err = p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	fmt.Println(r.URL.Path)
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid page title")
	}
	fmt.Println(m)
	return m[2], nil
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

var validPath = regexp.MustCompile("^/(edit|view|save)/([a-zA-Z0-9]+)$")

func main() {
	http.HandleFunc("/view/", viewhandler)
	http.HandleFunc("/edit/", edithandler)
	http.HandleFunc("/save/", savehandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
