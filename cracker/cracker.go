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

type Led struct {
	Name string
	Port string
	State int
	QueuedState int
}

var log 	*logger.Logger
var dbh 	*DBHandler
var ledList	[]Led


/**
 * Starts up the main thread.
 * 
 * @name main
 */
func main () {
	ledList = []Led{}
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

    /* Adding Leds to a list */
    ledList = append(ledList, Led{Name: "internet_access", Port: "GPIO14", State: LEDOFF, QueuedState: LEDON})


	log.Info("Sleeping...")
	time.Sleep(10 * time.Second)
}


/**
 * Controls the LED, that will indicate the status of the cracker.
 * 
 * @name blickLight
 */
func LEDController(led Led) {
	for true {
		for _,led := range ledList {
			if led.QueuedState != led.State {
	        	led.State = led.QueuedState
			    switch led.State {
			        case LEDBLINK:
			            go LEDBlink(1, led)
			        case LEDON: 
			        	LEDBlink(0, led)
			        default:
			            LEDBlink(-1, led)
			    }
			}
		} 
		time.Sleep(100 * time.Millisecond)
	}
}

/**
 * Makes the LED blink NumBlinks times.
 * If NumBlinks = 0, the LED will only be turned on.
 * If NumBlinks = -1, the LED will only be turned off.
 * 
 * @name LEDBlink
 * @param NumBlinks
 */
func LEDBlink(NumBlinks int, led Led) {
	if NumBlinks == 0 {
		// Leave the LED on.
		return
	}
	if NumBlinks == -1 {
		// Leave the LED on.
		return 
	}
	for true {
		if led.State != LEDBLINK {
			return 
		}
		log.Info("Blinking..")
		time.Sleep(1 * time.Second)
	}
}


/**
 * Scans and updates the database file based on what 
 * files are found.
 * 
 * @name ScanUpdate
 */
func ScanUpdate() {
    var err error

    flist := scanDir(DIRWPA)
    for _,file := range flist {
        if strings.Contains(file.Name(), ".cap") {
            var bssid []byte
            bssidfile := DIRWPA + strings.Split(file.Name(), ".")[0] + ".bssid"
            if bssid, err = ioutil.ReadFile(bssidfile); err != nil {
                log.Error(fmt.Sprintf("Missing file (%s), will ignore: %s", bssidfile, file.Name()))
                continue
            }
            log.Info(fmt.Sprintf("Found bssid: %s", bssid))
            dbh.StoreWpa(&Wpa{name:file.Name(), bssid:string(bssid)})
        }
    }

	flist = scanDir(DIRWORDLIST)
    for _,file := range flist {
        dbh.StoreWordlist(&Wordlist{name:file.Name(), size:string(file.Size()), avg_run:0})
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
