package main

import (
	"html/template"
	"net/http"
	"log"
//	"strings"
)

const LOGPATH string = "page.log"

var templates = template.Must(template.ParseGlob("html/*"))

type Page struct {
	Title string
}


func homeHandler(w http.ResponseWriter, r *http.Request) {

	p := &Page{
		Title: "This is a test page!",
	}

    err := templates.ExecuteTemplate(w, "home", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}


func initlog() {
//	log.SetOutput(f) 
	log.SetFlags(log.Ltime|log.Lshortfile)
}

func main() {
	initlog()
	
	/* Enable fetching of files. */
	resources := http.FileServer(http.Dir("./recources/"))
    http.Handle("/recources/", http.StripPrefix("/recources/", resources))
	
	http.HandleFunc("/", homeHandler)
	
	log.Printf("|Running...")
	log.Fatal(http.ListenAndServe(":8080", nil)) 
}
