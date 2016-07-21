package main

import (
	"golang.org/x/net/html"
//	"fmt"
	"net/http"
	"strings"
	 "strconv"
)

type Image struct {
	Url string
}


type Parse struct {
	StartUrl string
	Links []string 
	Images []Image
}


func (this *Parse) init(StartUrl string) {
	this.StartUrl = StartUrl
	this.Links = []string{}
	this.Images = []Image{}
}

/*
 * Parse
 *
 * Loops a http response, and applies a given function on it.
 */
func (this *Parse) Parse(resp *http.Response, needle string, fn func(html.Token)) {
	z := html.NewTokenizer(resp.Body)

	for {
	    tt := z.Next()

	    switch {
	    case tt == html.ErrorToken:
	        return
	    case tt == html.StartTagToken:
	        t := z.Token()
	        
	        if isCatch := t.Data == needle; isCatch {
	        	fn(t)
	        }
	    }
	}	
}

///////////////////////////////// Parsers ////////////////////////////////////////////

/*
 * parser_fourchan_links
 *
 * Site: 4chan.org
 * Analises a html token in order to retrieve a link.
 */
func (this *Parse) parser_fourchan_links() (func(token html.Token)) {
	return func(token html.Token) {
    	var urlfound string = ""
    	var linkfound bool = false
        for _, attr := range token.Attr {
        	if strings.Contains(attr.Val, "replylink") {
        		linkfound = true
        	}
        	if strings.Contains(attr.Val, "thread") {
        		urlfound = attr.Val
        	}	            	
        }
        if linkfound {
        	if link := strings.Split(urlfound, "/"); len(link) > 2 {
        		this.Links = append(this.Links, this.StartUrl + urlfound)
        	}
        }
    }
}

/*
 * parser_images
 *
 * Site: General
 * Collects an image given a token.
 */
func (this *Parse) parser_images() (func(token html.Token)) {
	return func(token html.Token) {
		for _, attr := range token.Attr {
			if strings.Contains(attr.Key, "src") {
				this.Images = append(this.Images, Image{attr.Val})
			}
		}
	}
}

/*
 * parser_links
 *
 * Site: General
 * Collects a link given a token.
 */
func (this *Parse) parser_links() (func(token html.Token)) {
	// TODO: Remove dubbles
	// TODO: Add start url
	return func(token html.Token) {
		for _, attr := range token.Attr {
			if strings.Contains(attr.Key, "href") {
				link := attr.Val
				if ! strings.Contains(link, "htt") {
					link = this.StartUrl + link
				}
				if ! this.link_exists(link) {
					this.Links = append(this.Links, link)
				}
			}
		}		
	}
}



///////////////////////////////// Special ////////////////////////////////////////////

/*
 * fourchan_img_clean
 *
 * Site: 4chan
 * Cleans image urls.
 */
func (this *Parse) fourchan_img_clean() {
	var images []Image
	for _,img := range this.Images {
		img.Url = strings.Replace(img.Url, "s", "", -1)
		images = append(images, img)
	}
	this.Images = images
}


func (this *Parse) get_wish_products(resp *http.Response) {
	
}


/*
 * avanza_get_sellprice
 *
 * Site: Avanza
 * Gets the current sellprice from a given httpResponse
 */
func (this *Parse) avanza_get_sellprice(resp *http.Response) float64 {
	z := html.NewTokenizer(resp.Body)

	for {
	    tt := z.Next()

	    switch {
	    case tt == html.ErrorToken:
	        return 0.0
	    case tt == html.StartTagToken:
	        t := z.Token()
	        
	        if isCatch := t.Data == "span"; isCatch {
				for _, attr := range t.Attr {
					if strings.Contains(attr.Val, "sellPrice") {
						z.Next();
						tt := z.Token();
						strval := strings.Replace(tt.String(), ",", ".", -1);
						value, _ := strconv.ParseFloat(strval, 64)
						return value;
					}
				}	
	        }
	    }
	}	
}

///////////////////////////////// Help Functions ////////////////////////////////////////////

func (this *Parse) link_exists(link string) bool {
	for _, l := range this.Links {
		if l == link {
			return true
		}
	}
	return false
}
