package main

import (
//	"html/template"
	"io/ioutil"
	"net/http"
	"fmt"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (string) {
	filename := "html/" + title + ".html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(body)
}

func renderTemplate(w http.ResponseWriter, page string) {
	fmt.Fprintf(w, page)
}

func handlerHome(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, loadPage("home"))
}

func main() {
	http.HandleFunc("/home/", handlerHome)
	http.HandleFunc("/", handlerHome)

	fmt.Println("Running...")
	http.ListenAndServe(":8080", nil)
}