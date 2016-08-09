package main

import (
    "github.com/apsdehal/go-logger"
    "os"
    "time"
//    "fmt"
//    "io/ioutil"
//    "flag"
//    "errors"
)

/* -- Constants -- */
const LEDOFF 	= 0
const LEDON 	= 1
const LEDBLINK 	= 2

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
