package main

import (
	"fmt"
	"io/ioutil"
	)

func main(){
	fmt.Println("hejsan")
	arr := read_file()
	for _, value := range arr {
		str := string(value)
		fmt.Println(str)
	}

}


func read_file() []byte{
	str, _ := ioutil.ReadFile("input_day1")
	return str
}
