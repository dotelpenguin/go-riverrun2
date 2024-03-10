// RiverRun AKA Pitfall Clone

// Playfield Array Legend
// 0 - Blank
// 1 - Wall 1
// 2 - Wall 2
// 3 - Wall 3
// 4 - Water 1 (Random Percentage from theme)
// 5 - Obstacle 1
// 6 - Obstacle 2
// 7 - Bonus 1
// 8 - Bonus 2
// 9 - Finish

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
)

// Declare a global variables

var playerXpos int = 10 // Player Initial X Position
var playerYpos int = 1  // Player Initial Y Position

var playfieldXoffset int = 19 // Playfield X Offset
var playfieldYoffset int = 2  // Playfield Y Offset

var playfieldArray [20][20]int // 20x20 playfield array

var playfieldTheme [10][10]string // 10x10 playfield theme array TODO: Implement this

func main() {
	initializeGame()
	// Menu Here
	gameRun()
}

func gameRun() { // Main function LOOP
	// Initialize the screen
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e := s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	s.SetStyle(defStyle)

	playfieldBoxes(s)

	for { // Game Loop
		playfieldDisplay(s)
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				s.Fini()
				os.Exit(0)
			case tcell.KeyRight:
				playerXpos++
				if gameCheckBoundries() { // Check if player is out of bounds
					playerXpos--
				}
				s.Show()
			case tcell.KeyLeft:
				playerXpos--
				if gameCheckBoundries() { // Check if player is out of bounds
					playerXpos++
				}
				s.Show()
			case tcell.KeyUp:
				playerYpos--
				if gameCheckBoundries() { // Check if player is out of bounds
					playerYpos++
				}
				s.Show()
			case tcell.KeyDown:
				playerYpos++
				if gameCheckBoundries() { // Check if player is out of bounds
					playerYpos--
				}
				s.Show()
			}
			style := tcell.StyleDefault.Foreground(tcell.ColorWhite.TrueColor()).Background(tcell.ColorBlack)
			collision, message := gameCheckCollision()
			debugStr1 := fmt.Sprintf("Key: %d", ev.Key())
			debugStr2 := fmt.Sprintf("Player: %d x %d  Array: ? x ?", playerXpos, playerYpos)
			debugStr3 := fmt.Sprintf("Collision: %s, %s", message, strconv.FormatBool(collision))
			printStr(s, playfieldXoffset+25, playfieldYoffset+1, style, debugStr1)
			printStr(s, playfieldXoffset+25, playfieldYoffset+2, style, debugStr2)
			printStr(s, playfieldXoffset+25, playfieldYoffset+3, style, debugStr3)
		}
	}
}

func initializeGame() {
	fmt.Println("Checking for game initialization.") // Check game requirements
	fmt.Println("Checking terminal size...........") // Check terminal size
	fmt.Println("Checking terminal encoding.......") // Check terminal encoding
	fmt.Println("Checking terminal speed..........") // Check terminal speed
	fmt.Println("Loading game assets..............") // Load game assets and check for errors
	fmt.Println("Loading dropfile.................") // Load dropfile and check for errors door.sys, chain.txt, door32.sys, doorinfo.def
	fmt.Println("Checking for other players.......") // Check for other players socket? pipe? file?
	fmt.Println("Initializing game................") // Initialize game
	playfieldBuild()                                 // Build the initial playfield array

	time.Sleep(2 * time.Second)
}
