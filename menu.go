package main

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

func menuDisplay(s tcell.Screen) {
	//playfieldDisplay(s)
	style := tcell.StyleDefault.Foreground(tcell.ColorYellow.TrueColor()).Background(tcell.ColorBlack)

	printStr(s, playfieldXoffset+1, playfieldYoffset+1, style, "==================")
	printStr(s, playfieldXoffset+8, playfieldYoffset+2, style, "Menu")
	printStr(s, playfieldXoffset+1, playfieldYoffset+3, style, "==================")
	printStr(s, playfieldXoffset+2, playfieldYoffset+5, style, "1: Start Game")
	printStr(s, playfieldXoffset+2, playfieldYoffset+6, style, "2: Game Mode")
	printStr(s, playfieldXoffset+2, playfieldYoffset+7, style, "3: Game Level")
	printStr(s, playfieldXoffset+2, playfieldYoffset+8, style, "4: Game Speed")
	printStr(s, playfieldXoffset+2, playfieldYoffset+9, style, "Q: Quit")

	playfieldUpdateStatus(s)
	s.Show()

	for {
		switch mev := s.PollEvent().(type) {
		case *tcell.EventKey:
			switch mev.Key() {
			case tcell.KeyRune:
				switch mev.Rune() {
				case '1':
					gameStart = true
					gameScore = 0
					gameRun(s)
				case 'q':
					s.Fini()
					os.Exit(0)
				default:
					continue
				}
			}
		}
	}
}
