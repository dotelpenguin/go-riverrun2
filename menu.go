package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
)

func menuDisplay(s tcell.Screen) {
	gameStart = false

	style := tcell.StyleDefault.Foreground(tcell.ColorYellow.TrueColor()).Background(tcell.ColorBlack)

	var padMenu string = strings.Repeat(" ", 20)
	for i := 0; i < 20; i++ {
		printStr(s, playfieldXoffset, playfieldYoffset+i, style, padMenu)
	}

	printStr(s, playfieldXoffset+1, playfieldYoffset+1, style, "==================")
	printStr(s, playfieldXoffset+1, playfieldYoffset+2, style, "       Menu")
	printStr(s, playfieldXoffset+1, playfieldYoffset+3, style, "==================")
	printStr(s, playfieldXoffset+1, playfieldYoffset+5, style, "1: Start")
	printStr(s, playfieldXoffset+1, playfieldYoffset+6, style, "2: Mode")
	printStr(s, playfieldXoffset+1, playfieldYoffset+7, style, "3: Level")
	printStr(s, playfieldXoffset+1, playfieldYoffset+8, style, "4: Speed")
	printStr(s, playfieldXoffset+1, playfieldYoffset+9, style, "5: Debug: "+strconv.FormatBool(debug))
	printStr(s, playfieldXoffset+1, playfieldYoffset+10, style, "Q: Quit")

	playfieldUpdateStatus(s)

	s.Sync()

	for {
		s.Show()
		switch mev := s.PollEvent().(type) {
		case *tcell.EventKey:
			switch mev.Key() {
			case tcell.KeyRune:
				switch mev.Rune() {
				case '1':
					gameStart = true
					gameScore = 0
					gameRun(s)
					return
				case '2':
					gameMode++
					if gameMode > 5 {
						gameMode = 0
					}
					return
				case '3':
					gameLevel++
					if gameLevel > 99 {
						gameLevel = 0
					}
					return
				case '4':
					gameSpeed++
					if gameSpeed > 200 {
						gameSpeed = 0
					}
					return
				case 'q':
					s.Fini()
					os.Exit(0)
				case '5':
					debug = !debug
					return
				default:
					continue
				}
			}
		}
	}
}
