package main
/*********************************************
 *	- Read config from file.
 *	
 *	
 *
 *
 *
 ********************************************/

import (
	"net/http"
	"fmt"
	"time"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"sync"
)


type Config struct {
	StartUrl string 
}

var images []*Image
var dbh *DBHandler
var parse *Parse
var config *Config
var wg sync.WaitGroup


func main() {
	if err := loadConfig(); err != nil {
		panic("FAIL!, Failed to load config!")
	}
	dbh = new(DBHandler)
	dbh.init()
	defer dbh.Close()

	parse = new(Parse)
	parse.init(config.StartUrl)

	get_links(config.StartUrl)

	fmt.Println("Getting images...")
	for _,link := range parse.Links {
		wg.Add(1)
		go findImges(link, true)
		fmt.Printf(".")
	} 
	fmt.Println("Done!")
	fmt.Println("Wating for goroutines...")
	wg.Wait()

	parse.fourchan_img_clean()
	for _,i := range parse.Images {
		fmt.Println(i.Url)
	}
	fmt.Printf("Totaly found %d images!\n", len(parse.Images))
}

/*
 * get_links
 *
 * Gets all the links on the given page
 */
func get_links(url string) {
	for i := 1; i < 11; i++ {
		resp, _ := http.Get(config.StartUrl + strconv.Itoa(i) + "/")
		parse.Parse(resp, "a", parse.parser_fourchan_links())
	}
}

func loadConfig() error {
	b, _ := ioutil.ReadFile("config.json")
	return json.Unmarshal(b, &config)
}

func findImges(url string, async bool) {
	if async {
		defer wg.Done()
	}
	resp, _ := http.Get(url)
	parse.Parse(resp, "img", parse.parser_images())
}



///////////////////////////////// WORK IN PROGRESS ////////////////////////////////////////////


	// list, err := dbh.Resource_geta()
	// if err != nil {
	// 	fmt.Printf("ERROR %s", err)
	// }

	// fmt.Printf("Got %d rows!\n", len(list))
	// for itr := range list {
	// 	fmt.Printf("Url: %s, Created: %d\n", list[itr].Url, list[itr].Created)
	// }




func add_resource(url string, time int) {
	resource := new(Resource)
	resource.Url = url
	resource.Created = time
	dbh.Resource_save(resource)
}


func routine() {
	add_resource("sdfsdfsdf", 55)

	list, err := dbh.Resource_geta()
	if err != nil {
		fmt.Printf("ERROR %s", err)
	}

	fmt.Printf("Got %d rows!\n", len(list))
	for itr := range list {
		fmt.Printf("Url: %s, Created: %d\n", list[itr].Url, list[itr].Created)
	}
	time.Sleep(5)
}