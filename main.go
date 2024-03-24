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
	"math/rand"
	"os"
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

var riverLeftboundry int = 5   // River Left Boundry
var riverRightboundry int = 18 // River Right Boundry
var riverWidth int = 10        // River Width
var riverMinwidth int = 12     // River Minimum Width
var riverMaxwidth int = 16     // River Maximum Width

var padLeftriver string = ""                // Pad Left Wall
var padRightriver string = ""               // Pad Right Wall
var padRiver string = ""                    // Pad River
var padLine string = "11111111111111111111" // Pad Line
var padLength int = 0                       // Pad Length

var gameMode int = 0       // Game Mode
var gameStart bool = false // Game Run
var gameLevel int = 0      // Game Level
var gameSpeed int = 1      // Game Speed
var gameSeed int64 = 0     // Game Seed
var gameScore int = 0      // Game Score
var gamescoreAlgo int = 0  // Game Last Score

var debug bool = false // Debug mode

var quit = make(chan bool)

func main() {

	// Initialize the game
	initializeGame()

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

	go gameLoop(s) // Start the background game loop
	playfieldBoxes(s)

	for {
		menuDisplay(s)
	}
}

func gameRun(s tcell.Screen) { // Main function LOOP

	playfieldBoxes(s)

	for { // Game Loop
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				gameStart = false
				menuDisplay(s)
			case tcell.KeyRight:
				playerXpos++
				if gameCheckBoundries() { // Check if player is out of bounds
					playerXpos--
				}
				playfieldDisplay(s)
			case tcell.KeyLeft:
				playerXpos--
				if gameCheckBoundries() { // Check if player is out of bounds
					playerXpos++
				}
				playfieldDisplay(s)
			case tcell.KeyUp:
				playerYpos--
				if gameCheckBoundries() { // Check if player is out of bounds
					playerYpos++
				}
				playfieldDisplay(s)
			case tcell.KeyDown:
				playerYpos++
				if gameCheckBoundries() { // Check if player is out of bounds
					playerYpos--
				}
				playfieldDisplay(s)
			default:
				continue
			}
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
	rand.Seed(time.Now().UnixNano())                 // Seed the random number generator TODO: Is this settable? Should it be?
	playfieldBuild()                                 // Build the initial playfield array
	time.Sleep(1 * time.Second)
}
