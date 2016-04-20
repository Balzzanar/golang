package main

import (
	"html/template"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)

const LOGPATH string = "page.log"
var templates = template.Must(template.ParseGlob("html/*"))

type Ingredience struct {
	Name string
	Amount string
}

type Recept struct {
	ID string
	Name string
	Ingrediences []Ingredience 
	Making []string 
	Notes []string
	Image string
	Description string
	Type string
}

type Page struct {
	Title string
	Recepts []Recept 
}

/**
 * loadRecepts
 *
 * Reads recepies from json formated files in a given dir.
 */
func loadRecepts() []Recept {
	var recepts []Recept
	files, err := ioutil.ReadDir("recept-files")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		b, _ := ioutil.ReadFile("recept-files/" + file.Name())
		recepts = append(recepts, parseRecept(b))
	}
	return recepts
}

/**
 * parseRecept
 *
 * Help function for loadRecepts, that parses the json recept.
 */
func parseRecept(raw []byte) Recept {
	var recept = Recept{}
	err := json.Unmarshal(raw, &recept)
	if err != nil {
		log.Println("error:", err)
	}
	return recept
}


func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "Sphann, home page!",
		Recepts: loadRecepts(),
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
	log.Fatal(http.ListenAndServe(":80", nil)) 
}
