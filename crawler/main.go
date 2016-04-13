package main

import (
	"fmt"
	"time"
)

var dbh *DBHandler

func main(){
	dbh = new(DBHandler)
	dbh.init()
	//defer dbh.Close()

	routine()
	routine()

	list, err := dbh.Resource_geta()
	if err != nil {
		fmt.Printf("ERROR %s", err)
	}

	fmt.Printf("Got %d rows!\n", len(list))
	for itr := range list {
		fmt.Printf("Url: %s, Created: %d\n", list[itr].Url, list[itr].Created)
	}
}

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
