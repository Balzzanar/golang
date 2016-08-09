package main

import (
    _ "github.com/mattn/go-sqlite3"
    "database/sql"
    "fmt"
)

/* -- Constants -- */
const TABLE_WPA = `create table if not exists wpa (id integer not null primary key, name text, bssid varchar(30));`
const TABLE_WORDLISTS = `create table if not exists wordlists (id integer not null primary key, name text, size varchar(10), avg_run int);`
const TABLE_RUNS = `create table if not exists runs (id_wpa int, id_wordlist, result text, time int, started int);`

type Wordlist struct {
	id int
	name string
	size string
	avg_run int
}



type DBHandler struct {
	db *sql.DB
}

/**
 * Opens a connection to the databasefile, creates one if it does not exits
 * 
 * @name Init
 */
func (this *DBHandler) Init() {
	var derr error
	this.db, derr = sql.Open("sqlite3", "./foo.db")
	if derr != nil {
		fmt.Println(derr)
	}
	this.createNewTable(TABLE_WPA)
	this.createNewTable(TABLE_WORDLISTS)
	this.createNewTable(TABLE_RUNS)
}


/**
 * Closes the connection to the databasefile
 * 
 * @name Close
 */
func (this *DBHandler) Close() {
	this.db.Close()
}


/**
 * Stores a wordlist to the databasefile
 * 
 * @name StoreWordlist
 */
func (this *DBHandler) StoreWordlist() {

}


/**
 * Runs a table script on the database file.
 * 
 * @name createNewTable
 */
func (this *DBHandler) createNewTable(tablescript string) {
	_, err := this.db.Exec(tablescript)
	if err != nil {
		log.Error(fmt.Sprintf("%q: %s\n", err, tablescript))
		return
	}
}

