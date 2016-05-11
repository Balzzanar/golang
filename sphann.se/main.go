package main

import (
	"html/template"
	"net/http"
	"strings"
	"log"
	"io/ioutil"
	"encoding/json"
)

const LOGPATH string = "page.log"
var templates = template.Must(template.ParseGlob("html/*"))
var EIDTIPADDR string = ""

type Ingredience struct {
	Name string
	Amount string
}

type Wine struct {
	Name string
	Number string
	GoodBad string
	Notes string
	Link string
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
	Wines []Wine
}

type Files struct {
	Name string
	Content string
}

type Page struct {
	Title string
	Recepts []Recept 
	Files []Files
	Edit bool
}


/**
 * loadRecepts
 *
 * Reads recepies from json formated files in a given dir.
 */
func loadRecepts() ([]Recept, []Files) {
	var recepts []Recept
	var filesarr []Files
	files, err := ioutil.ReadDir("recept-files")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		b, _ := ioutil.ReadFile("recept-files/" + file.Name())
		recepts = append(recepts, parseRecept(b))
		f := Files{Name: file.Name(), Content: string(b)}
		filesarr = append(filesarr, f)
	}
	return recepts, filesarr
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
	addr := r.RemoteAddr
	addr = strings.Split(addr, ":")[0]
	log.Println("Addr: " + addr+ "jjj: "+ EIDTIPADDR)
	edit := true
	if addr != EIDTIPADDR {
		edit = false
	}

	recerpts, files := loadRecepts()
	p := &Page{
		Title: "Sphann, home page!",
		Recepts: recerpts,
		Files: files,
		Edit: edit,
	}

    err := templates.ExecuteTemplate(w, "home", p)
    if err != nil {
		log.Fatal(err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func saveReceptHandler(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var recept Recept
    err := decoder.Decode(&recept)
    if err != nil {
        log.Fatal(err)
        panic("Failed to decode json")
    }
    log.Println(recept.ID)
	fname := recept.ID + ".json"
	log.Println(recept.ID)

	toFile, err := json.Marshal(recept)
    if err != nil {
        log.Fatal(err)
        panic("Failed to encode json")
    }
	e := ioutil.WriteFile("recept-files/" + fname, toFile, 0644)
	if e != nil {
    	panic(e)
	}
    w.WriteHeader(200)
}


func initlog() {
//	log.SetOutput(f) 
	log.SetFlags(log.Ltime|log.Lshortfile)
}

func loadConfig() {
	b, _ := ioutil.ReadFile("config/allowed_ip")
	EIDTIPADDR = strings.TrimSpace(string(b))
}

func main() {
	initlog()
	loadConfig()
	
	/* Enable fetching of files. */
	resources := http.FileServer(http.Dir("./recources/"))
    http.Handle("/recources/", http.StripPrefix("/recources/", resources))
	
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/saverecept", saveReceptHandler)

	log.Printf("|Running...")
	log.Fatal(http.ListenAndServe(":8080", nil)) 
}
