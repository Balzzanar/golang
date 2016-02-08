package main

import (
	"fmt"
	"io/ioutil"
	)

var floor int = 0
var char_pos int = 0
var first_char_pos int = 0

func main(){
	fmt.Println("----- Day 1 -----")
	arr := read_file()
	for _, value := range arr {
		char_pos += 1
		str := string(value)
		if str == "(" {
			floor += 1
		} else if str == ")"{
			floor -= 1
		}

		if floor == -1 && first_char_pos == 0 {
			first_char_pos = char_pos
		}
	}
	fmt.Printf("Floor: %d\n", floor)
	fmt.Printf("Char pos: %d\n", first_char_pos)
}

func read_file() []byte{
	str, _ := ioutil.ReadFile("input_day1")
	return str
}
