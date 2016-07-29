package main

//import "os"
import "io/ioutil"
import "fmt"
import "regexp"
import "reflect"
import "sync"
import "net/http"


var wg sync.WaitGroup
var useragent string = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:45.0) Gecko/20100101 Firefox/45.0"

func main() {
	get_http_data("https://www.reddit.com/")
	/*
    data, _ := ioutil.ReadAll(os.Stdin)
    wg.Add(1)
    go run_regexp(data);
    fmt.Printf("Waiting for go-routine to finnish\n")
	wg.Wait()
*/
}

func run_regexp(data []byte) bool{
	regexpStr := `<td>(\S+)</td>`
	defer wg.Done()

	re := regexp.MustCompile(regexpStr)
	result := re.FindAllStringSubmatch(string(data), -1);
	fmt.Printf("%q\n", result)
	fmt.Printf("\n%s\n", reflect.TypeOf(result))

	for _,value := range result {
		fmt.Printf("%s is of type: %s \n", value, reflect.TypeOf(value))
	}
	return true
}

func get_http_data(url string) {
    client := &http.Client{}

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println(err)
    }

    req.Header.Set("User-Agent", useragent)
	resp, err := client.Do(req)
	htmlData, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(htmlData))
}