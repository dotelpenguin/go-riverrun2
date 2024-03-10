// Copyright 2022 The TCell Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Playfield related functions
//
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

	"github.com/gdamore/tcell/v2"
)

// Declare a global variables

var playerXpos int = 10 // Player Initial X Position
var playerYpos int = 1  // Player Initial Y Position

var playfieldXoffset int = 20 // Playfield X Offset
var playfieldYoffset int = 2  // Playfield Y Offset

var playfieldArray [20][20]int // 20x20 playfield array

var playfieldTheme [10][10]string // 10x10 playfield theme array

func main() { // Main function LOOP

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

	playfieldBuild() // Build the initial playfield array
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
				s.Show()
			case tcell.KeyLeft:
				playerXpos--
				s.Show()
			case tcell.KeyUp:
				playerYpos--
				s.Show()
			case tcell.KeyDown:
				playerYpos++
				s.Show()
			}
			style := tcell.StyleDefault.Foreground(tcell.ColorWhite.TrueColor()).Background(tcell.ColorBlack)
			collision, message := gameCheckCollision()
			debugStr1 := fmt.Sprintf("Key: %d - %d x %d", ev.Key(), playerXpos, playerYpos)
			debugStr2 := fmt.Sprintf("Collision: %s, %s", message, strconv.FormatBool(collision))
			printStr(s, playfieldXoffset+25, playfieldYoffset+1, style, debugStr1)
			printStr(s, playfieldXoffset+25, playfieldYoffset+2, style, debugStr2)
		}
	}
}
