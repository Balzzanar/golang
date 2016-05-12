package main

import (
		_ "github.com/go-sql-driver/mysql"
        "database/sql"
        "log"
        "sync"
        )


type DBHandler struct {
	lock *sync.Mutex
	db *sql.DB
}


func (this *DBHandler) init(){
	var err error 
	this.lock = new (sync.Mutex)
	this.db, err = sql.Open("mysql", "root:kti26md3ynku@tcp(192.168.0.1:3306)/crawdb")
	if err != nil {
		log.Println(err)
	} 
}

func (this *DBHandler) Close () {
	this.db.Close()
}




type Resource struct {
	Resource_id int
	Url string
	Created int
}


func (this *DBHandler) Resource_geta() ([]*Resource, error) {
	this.lock.Lock()
	
	var resources []*Resource

	rows, err := this.db.Query("select * from resources")
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		this.lock.Unlock()
		return resources, err
	}
	defer rows.Close()

	for rows.Next() {
		result := new(Resource)
		err = rows.Scan(&result.Resource_id, &result.Url, &result.Created)
		if err != nil {
			log.Fatal(err)
		}
		resources = append(resources, result)
	}
	this.lock.Unlock()
	return resources, err
}

func (this *DBHandler) Resource_save(resource *Resource) (error) {
	this.lock.Lock()

	query := "insert into resources (url, created) values(?, ?)"
	_, err := this.db.Query(query, resource.Url, resource.Created)
	if err != nil {
		log.Fatal(err)
	}
	this.lock.Unlock()
	return err
}

