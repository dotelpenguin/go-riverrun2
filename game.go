package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

func gameCheckCollision() (bool, string) {
	switch playfieldArray[playerYpos][playerXpos] {
	case 1:
		return true, "Wall 1"
	case 2:
		return true, "Wall 2"
	case 3:
		return true, "Wall 3"
	case 4:
		return false, "Water 1"
	case 5:
		return true, "Obstacle 1"
	case 6:
		return true, "Obstacle 2"
	case 7:
		return true, "Bonus 1"
	case 8:
		return true, "Bonus 2"
	case 9:
		return true, "Finish"
	}
	return false, "None"
}

func gameCheckBoundries() bool {
	// Check if player is out of bounds
	if playerXpos < 0 || playerXpos > 19 || playerYpos < 0 || playerYpos > 19 {
		return true
	}
	return false
}

func gameLoop(s tcell.Screen) { // Background Game Loop
	for {
		select {
		case <-quit:
			return
		default:
			if gameStart {
				time.Sleep(time.Duration(200-(gamescoreAlgo*10)) * time.Millisecond) // Pause for 200 milliseconds - gamespeedalgo
				playfieldDisplay(s)
				collision, message := gameCheckCollision()
				if collision {
					//gameStart = false
					style := tcell.StyleDefault.Foreground(tcell.ColorYellow.TrueColor()).Background(tcell.ColorBlack)
					// Insert game over function here
					printStr(s, playfieldXoffset+2, playfieldYoffset+4, style, "Message: "+message)
					s.Show()
				}
				gameAdvance()
				playfieldUpdateStatus(s)
			}
			if debug {
				playfieldDebug(s)
			}
		}
	}
}

func gameAdvance() { // Advance the game steps + adjust difficulty
	gameScore++
	gamescoreAlgo = gameScore / 100
	riverMinwidth = 12 - gamescoreAlgo
	riverMaxwidth = 16 - gamescoreAlgo
	if riverMinwidth < 3 {
		riverMinwidth = 3
		riverMaxwidth = 5
	}
	if gamescoreAlgo > 20 {
		gamescoreAlgo = 20
	}
	playfieldGenerateNewLine()
}
