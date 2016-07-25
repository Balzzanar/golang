package main

import "os"
import "io/ioutil"
import "fmt"
import "regexp"
import "reflect"

func main() {
    data, _ := ioutil.ReadAll(os.Stdin)
	re := regexp.MustCompile(`<td>(\S+)</td>`)
	result := re.FindAllStringSubmatch(string(data), -1);
	fmt.Printf("%q\n", result)
	fmt.Printf("\n%s\n", reflect.TypeOf(result))

	for _,value := range result {
		fmt.Printf("%s is of type: %s \n", value, reflect.TypeOf(value))
	}
}
