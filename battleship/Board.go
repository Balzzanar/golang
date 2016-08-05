package main

import (
	"fmt"
)

type Coordinate struct {
	XCord int
	YCord int
}

type Board struct {
	board [10][10]string
	markerShip string
	markerHit string
	markerMiss string
	markerDestroyd string
	markerBlank string	
}

/**  
 * Initis the board.
 * Fills it with ships at 'known' locations, if the DUMMYBOARD was set.
 *
 * @name Init
 */
func (this *Board) Init() {
	// Set the markers
	this.markerShip = "S"
	this.markerHit = "H"
	this.markerMiss = "M"
	this.markerDestroyd = "D"
	this.markerBlank = "-"
	
	this.board = [10][10]string{}
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			this.board[x][y] = this.markerBlank
		}
	}

	if *DUMMYBOARD {	// Used for loading the board with ships at the dummy location.
		this.addDummyShips()
		log.Info("Dummy board loaded...")
	} else {
		log.Info("Board loaded...")
	}
}


/**  
 * Place a marker on the board.
 *
 * @name PlaceMarker
 * @param marker
 * @param coordinate
 */
func (this *Board) PlaceMarker(marker string, coordinate *Coordinate) {
	// TODO: Implement checks, so a marker can't be placed other markers. 
	//		ie. miss on destroyed.
	if (this.board[coordinate.XCord][coordinate.YCord] != this.markerBlank){
		panic("aboo")
	}
	this.board[coordinate.XCord][coordinate.YCord] = marker	
}


/**  
 * Prints the board.
 *
 * @name print
 */
func (this *Board) Print() {
	fmt.Printf("The Board:\n")
	fmt.Printf("  ")
	for x := 0; x < 10; x++ {
		fmt.Printf("%d ", x)
	}
	fmt.Printf("\n")
	for y := 0; y < 10; y++ {
		fmt.Printf("%d ", y)
		for x := 0; x < 10; x++ {
			fmt.Printf(this.board[y][x])
			fmt.Printf(" ")
		}
		fmt.Printf("\n")
	}
}

/**  
 * Places some new ships 
 *
 * @name addDummyShips
 */
func (this *Board) addDummyShips() {
	this.board[5][5] = this.markerShip
	this.board[6][5] = this.markerShip
	this.board[7][5] = this.markerShip
	this.board[8][5] = this.markerShip


	this.board[2][7] = this.markerShip
	this.board[3][7] = this.markerShip

	this.board[2][3] = this.markerShip
	this.board[2][4] = this.markerShip
	this.board[2][5] = this.markerShip

	this.board[7][2] = this.markerShip
	this.board[7][3] = this.markerShip
}
