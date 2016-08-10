package main

import (
    "github.com/apsdehal/go-logger"
    "os"
    "time"
    "fmt"
    "strings"
    "io/ioutil"
//    "flag"
//    "errors"
)

/* -- Constants -- */
const LEDOFF 		= 0
const LEDON 		= 1
const LEDBLINK 		= 2
const DIRWORDLIST	= "wordlist/"
const DIRWPA		= "tocrack/"

var log *logger.Logger
var LEDStatus int  /*  0 - off, 1 - on, 2 - blinking  */
var dbh *DBHandler


/**
 * Starts up the main thread.
 * 
 * @name main
 */
func main () {
    var err error
    log, err = logger.New("cracker", 1, os.Stdout)
    if err != nil {
        panic(err) 
    }
    dbh = new(DBHandler)
    dbh.Init()
    defer dbh.Close()

    dbh.StoreWordlist(&Wordlist{id:0, name:"rockyou.txt", size:"143MB", avg_run:30012313})
    dbh.GetAllWpa()
    ScanUpdate()
    LEDController()
}


/**
 * Controls the LED, that will indicate the status of the cracker.
 * 
 * @name blickLight
 */
func LEDController() {
	LEDStatus = LEDON
	for true {
	    switch LEDStatus {
	        case LEDBLINK:
	            LEDBlink(1)
	        case LEDON: 
	        	LEDBlink(0)
	        default:
	            LEDBlink(0)
	    }

		// Sleep 
		log.Info("Sleeping...")
		time.Sleep(5 * time.Second)
	}
}

/**
 * Makes the LED blink NumBlinks times.
 * If NumBlinks = 0, the LED will only be turned on.
 * 
 * @name LEDBlink
 * @param NumBlinks
 */
func LEDBlink(NumBlinks int) {
	if NumBlinks == 0 {
		// Leave the LED on.
	}
	log.Info("Blinking..")
	time.Sleep(1 * time.Second)
}


/**
 * Scans and updates the database file based on what 
 * files are found.
 * 
 * @name ScanUpdate
 */
func ScanUpdate() {
    flist := scanDir(DIRWPA)
    var err error
    for _,file := range flist {
        if strings.Contains(file.Name(), ".cap") {
            var bssid []byte
            bssidfile := DIRWPA + strings.Split(file.Name(), ".")[0] + ".bssid"
            if bssid, err = ioutil.ReadFile(bssidfile); err != nil {
                log.Error(fmt.Sprintf("Missing file: %s", bssidfile))
            }
            log.Info(fmt.Sprintf("Found bssid: %s", bssid))
            dbh.StoreWpa(&Wpa{name:file.Name(), bssid:string(bssid)})
        }
    }
}


/**
 * Scans a given directory and returns a list of 
 * all the files.
 * 
 * @name scanDir
 */
func scanDir(dir string) []os.FileInfo {
    files := []os.FileInfo{}
    d, err := os.Open(dir)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer d.Close()
    fi, err := d.Readdir(-1)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    for _, fi := range fi {
        if fi.Mode().IsRegular() {
            files = append(files, fi)
        }
    }
    return files
}