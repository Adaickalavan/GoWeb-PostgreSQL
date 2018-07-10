package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
)

func viewhandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func edithandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func savehandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

var validPath = regexp.MustCompile("^/(edit|view|save)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	fmt.Println("inside makehandler")
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("inside func 2")
		m := validPath.FindStringSubmatch(r.URL.Path)
		fmt.Println("after find substring")
		if m == nil {
			http.NotFound(w, r)
			fmt.Println("inside m nil")
			return
		}
		fmt.Println("outside m nil")
		fn(w, r, m[2])
	}
}

func main2() {
	http.HandleFunc("/view/", makeHandler(viewhandler))
	http.HandleFunc("/edit/", makeHandler(edithandler))
	http.HandleFunc("/save/", makeHandler(savehandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
