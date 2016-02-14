package main

import (
		"time"
        "fmt"
        )

func main(){

	for true {
		fmt.Println("Running another loop")
		time.Sleep(10000 * time.Millisecond)
	}
}

