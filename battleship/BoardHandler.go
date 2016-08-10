/**  
 * BoardHandler 
 *
 *   
 */
package main

import (
    "github.com/apsdehal/go-logger"
    "os"
    "fmt"
    "flag"
 //   "reflect"
 //   "errors"
)

var log *logger.Logger
var board *Board
var DUMMYBOARD *bool

func main() {
	var err error
    log, err = logger.New("log", 1, os.Stdout)
    if err != nil {
        panic(err) 
    }
    if err = init_args(); err != nil {
    	panic(fmt.Sprintf("Bad argument input! (%s)\nuse '--help' to get a list of arguments", err))
    }
    log.Info("BoardHandler Started...")

    board = new(Board)
    board.Init()
    log.Info("Entering the main gameloop")
    GameLoop()
}


/**  
 * Starts the main Game loop, in which it will remain until
 * the game ends.
 *
 *  - Will recover from panics and weain on the user for bad input. 
 *
 * @name GameLoop
 */
func GameLoop() {
    defer func() {
        if r := recover(); r != nil {
            reason := panicHandler(r)
            log.Error(fmt.Sprintf("Error! (%s)", reason))
            GameLoop()
        }
    }()

    board.Print()
    for true {
        fmt.Printf("Insert Cords (X Y): ")
        cord := new(Coordinate)
        fmt.Scanf("%d %d", &cord.XCord, &cord.YCord)
        board.PlaceMarker(board.markerMiss, cord)
        board.Print()        
    }
}


/**  
 * Takes care of all the errors that occurs in the GameLoop
 *
 * @name panicHandler
 * @return string
 */
func panicHandler(e interface{}) string {
    var err error
    switch x := e.(type) {
        case error:
            err = x
        default:
            err = fmt.Errorf("%v", x)
    }

    switch (err.String()) {
        case "aboo":
            fmt.Println("OMOGMG")
        default:
            fmt.Println("Bad input!")
    }

    return ""   
}

/**  
 * Loads the arguments sent to the BoardHandler 
 *
 * @name init_args
 * @return error
 */
func init_args() error {
	var result error 
	DUMMYBOARD = flag.Bool("d", false, "Use a dummy board")
	flag.Parse()
	return result	
}

