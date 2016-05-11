package main

import (
	"golang.org/x/net/html"
	"net/http"
	"fmt"
	"time"
	"strings"
	"strconv"
	//"io/ioutil"
)

var dbh *DBHandler

func main(){

	/*********************************************
	 *	- Read config from file.
	 *	
	 *	
	 *
	 *
	 *
	 ********************************************/

	dbh = new(DBHandler)
	dbh.init()
	defer dbh.Close()

	url := "http://boards.4chan.org/diy/"

	var urls []string 
	for i := 1; i < 11; i++ {
		urls = append(urls, get_links(url + strconv.Itoa(i) + "/", url)...)
	}

	for _,url := range urls {
		fmt.Println(url)
	}

	// list, err := dbh.Resource_geta()
	// if err != nil {
	// 	fmt.Printf("ERROR %s", err)
	// }

	// fmt.Printf("Got %d rows!\n", len(list))
	// for itr := range list {
	// 	fmt.Printf("Url: %s, Created: %d\n", list[itr].Url, list[itr].Created)
	// }
}

func add_resource(url string, time int) {
	resource := new(Resource)
	resource.Url = url
	resource.Created = time
	dbh.Resource_save(resource)
}


func get_links(url string, orgiurl string) []string {
	var urls []string 
	resp, _ := http.Get(url)
	z := html.NewTokenizer(resp.Body)

	for {
	    tt := z.Next()

	    switch {
	    case tt == html.ErrorToken:
	        // End of the document, we're done
	        return urls
	    case tt == html.StartTagToken:
	        t := z.Token()

	        isAnchor := t.Data == "a"

	        if isAnchor {
	        	var urlfound string
	        	var linkfound bool = false
	            for _, attr := range t.Attr {
	            	if strings.Contains(attr.Val, "replylink") {
	            		linkfound = true
	            	}
	            	if strings.Contains(attr.Val, "thread") {
	            		urlfound = attr.Val
	            	}	            	
	            }
	            if linkfound {
	            	urls = append(urls, orgiurl + urlfound)
	            }
	        }
	    }
	}	
	return urls
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
