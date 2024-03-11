package main

import (
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

func playfieldBoxes(s tcell.Screen) {

	// Debug Ruler for 40/80
	style := tcell.StyleDefault.Foreground(tcell.ColorRed.TrueColor()).Background(tcell.ColorBlack)
	var pad80 string = strings.Repeat(":", 80)
	printStr(s, 0, 0, style, pad80)
	printStr(s, 55, 0, style, "PC 80 Field")

	style = tcell.StyleDefault.Foreground(tcell.ColorYellow.TrueColor()).Background(tcell.ColorBlack)
	var pad40 string = strings.Repeat(".", 40)
	printStr(s, 0, 0, style, pad40)
	printStr(s, 15, 0, style, "C64 40 Field")
	// End Debug Ruler

	// Draw the playfield box
	drawBox(s, playfieldXoffset-1, playfieldYoffset-1, playfieldXoffset+20, playfieldYoffset+20, style)

	// Draw the Score box
	drawBox(s, 0, playfieldYoffset-1, 16, playfieldYoffset+4, style)

	// Draw the Info box
	drawBox(s, 0, playfieldYoffset+5, 16, playfieldYoffset+20, style)

	// Draw the debug/artwork box. Disabled in 40 column mode
	drawBox(s, playfieldXoffset+22, playfieldYoffset-1, 79, playfieldYoffset+20, style)
}

func playfieldDebug(s tcell.Screen) {
	// Print the debug info
	style := tcell.StyleDefault.Foreground(tcell.ColorYellow.TrueColor()).Background(tcell.ColorBlack)
	drawBox(s, playfieldXoffset+22, playfieldYoffset-1, 79, playfieldYoffset+20, style)

	// Lets find a better way to do this
	var pad35 string = strings.Repeat(" ", 35)
	printStr(s, playfieldXoffset+25, playfieldYoffset+2, style, pad35)
	printStr(s, playfieldXoffset+25, playfieldYoffset+3, style, pad35)
	printStr(s, playfieldXoffset+25, playfieldYoffset+4, style, pad35)
	printStr(s, playfieldXoffset+25, playfieldYoffset+5, style, pad35)
	printStr(s, playfieldXoffset+25, playfieldYoffset+6, style, pad35)
	printStr(s, playfieldXoffset+25, playfieldYoffset+7, style, pad35)
	// End of better way

	collision, message := gameCheckCollision()
	printStr(s, playfieldXoffset+25, playfieldYoffset+1, style, "Debug Info")
	printStr(s, playfieldXoffset+25, playfieldYoffset+2, style, "Array Element: "+strconv.Itoa(playfieldArray[playerYpos][playerXpos]))
	printStr(s, playfieldXoffset+25, playfieldYoffset+3, style, "Collision: "+strconv.FormatBool(collision))
	printStr(s, playfieldXoffset+25, playfieldYoffset+4, style, "Message: "+message)
	printStr(s, playfieldXoffset+25, playfieldYoffset+5, style, "Player X: "+strconv.Itoa(playerXpos))
	printStr(s, playfieldXoffset+25, playfieldYoffset+6, style, "Player Y: "+strconv.Itoa(playerYpos))
	printStr(s, playfieldXoffset+25, playfieldYoffset+7, style, "Speed: "+strconv.Itoa(gameSpeed))
	printStr(s, playfieldXoffset+25, playfieldYoffset+8, style, "Level: "+strconv.Itoa(gameLevel))
	printStr(s, playfieldXoffset+25, playfieldYoffset+9, style, "Mode: "+strconv.Itoa(gameMode))
	printStr(s, playfieldXoffset+25, playfieldYoffset+10, style, "Start: "+strconv.FormatBool(gameStart))
}

func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style) {
	// Double check to see if there isn't a native function for this.
	// Draw the top and bottom
	for x := x1; x <= x2; x++ {
		s.SetContent(x, y1, '-', nil, style)
		s.SetContent(x, y2, '-', nil, style)
	}

	// Draw the sides
	for y := y1 + 1; y < y2; y++ {
		s.SetContent(x1, y, '|', nil, style)
		s.SetContent(x2, y, '|', nil, style)
	}

	// Draw the corners
	s.SetContent(x1, y1, '+', nil, style) // top left
	s.SetContent(x2, y1, '+', nil, style) // top right
	s.SetContent(x1, y2, '+', nil, style) // bottom left
	s.SetContent(x2, y2, '+', nil, style) // bottom right
}

func printStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

func playfieldBuild() { // Initial Playfield Array + Title Screen
	playfieldGrid := `
	12344444444444444321
	11234444444444444321
	11123444444444443211	
	11234444444444432111
	12344444444444443211
	11234444444444432111
	11123444474444321111
	11112344444443211111
	11111234444444321111
	11111123444544432111
	11111234444444443211
	11111234444444444321
	11111123484444444321
	11111111234444444321
	11111112344444443211
	11111123464444432111
	11111234444444321111
	11111123499943211111
	11111234444444321111
	11111123444444432111` // todo: implement theme/assets from file

	// Dirty remove of non Integers Characters
	playfieldGrid = strings.ReplaceAll(playfieldGrid, "\n", "")
	playfieldGrid = strings.ReplaceAll(playfieldGrid, " ", "")
	playfieldGrid = strings.ReplaceAll(playfieldGrid, "\t", "")

	index := 0
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			num, err := strconv.Atoi(string(playfieldGrid[index]))
			if err != nil {
				panic(err)
			}
			playfieldArray[i][j] = num
			index++
		}
	}
}

func playfieldDisplay(s tcell.Screen) {
	// Print the array
	colorCode := 0
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			colorCode = playfieldArray[i][j]
			// Set color based on colorCode. todo: implement theme/assets from file
			switch colorCode {
			case 0:
				style := tcell.StyleDefault.Foreground(tcell.ColorBlack.TrueColor()).Background(tcell.ColorBlack)
				s.SetContent(j+playfieldXoffset, i+playfieldYoffset, '.', nil, style) // black/black - Blank
			case 1:
				style := tcell.StyleDefault.Foreground(tcell.ColorYellow.TrueColor()).Background(tcell.ColorGreen)
				s.SetContent(j+playfieldXoffset, i+playfieldYoffset, ':', nil, style) // yellow/green - Wall 1
			case 2:
				style := tcell.StyleDefault.Foreground(tcell.ColorYellow.TrueColor()).Background(tcell.ColorGreen)
				s.SetContent(j+playfieldXoffset, i+playfieldYoffset, ']', nil, style) // yellow/green - Wall 2
			case 3:
				style := tcell.StyleDefault.Foreground(tcell.ColorYellow.TrueColor()).Background(tcell.ColorGreen)
				s.SetContent(j+playfieldXoffset, i+playfieldYoffset, '}', nil, style) // yellow/green - Wall 3
			case 4:
				style := tcell.StyleDefault.Foreground(tcell.ColorDarkCyan.TrueColor()).Background(tcell.ColorDarkBlue)
				s.SetContent(j+playfieldXoffset, i+playfieldYoffset, '~', nil, style) // cyan/blue - Water 1
			case 5:
				style := tcell.StyleDefault.Foreground(tcell.ColorRed.TrueColor()).Background(tcell.ColorDarkBlue)
				s.SetContent(j+playfieldXoffset, i+playfieldYoffset, 'X', nil, style) // red/blue - Obstacle 1
			case 6:
				style := tcell.StyleDefault.Foreground(tcell.ColorRed.TrueColor()).Background(tcell.ColorDarkBlue)
				s.SetContent(j+playfieldXoffset, i+playfieldYoffset, '^', nil, style) // red/blue - Obstacle 2
			case 7:
				style := tcell.StyleDefault.Foreground(tcell.ColorWhite.TrueColor()).Background(tcell.ColorDarkBlue)
				s.SetContent(j+playfieldXoffset, i+playfieldYoffset, '$', nil, style) // white/blue - Bonus 1
			case 8:
				style := tcell.StyleDefault.Foreground(tcell.ColorWhite.TrueColor()).Background(tcell.ColorDarkBlue)
				s.SetContent(j+playfieldXoffset, i+playfieldYoffset, '+', nil, style) // white/blue - Bonus 2
			case 9:
				style := tcell.StyleDefault.Foreground(tcell.ColorDarkMagenta.TrueColor()).Background(tcell.ColorDarkBlue)
				s.SetContent(j+playfieldXoffset, i+playfieldYoffset, '!', nil, style) // green - Finish
			default:
				style := tcell.StyleDefault.Foreground(tcell.ColorYellow.TrueColor()).Background(tcell.ColorBlack)
				s.SetContent(j+playfieldXoffset, i+playfieldYoffset, '.', nil, style) // reset to default color
			}
			style := tcell.StyleDefault.Foreground(tcell.ColorWhite.TrueColor()).Background(tcell.ColorDarkBlue)
			s.SetContent(playerXpos+playfieldXoffset, playerYpos+playfieldYoffset, '#', nil, style) // yellow/green - Wall 1
			s.Show()
		}

	}
}

func playfieldGenerateNewLine() { // Generates new line for the playfield and moves everything up
	for i := 1; i < 20; i++ {
		for j := 0; j < 20; j++ {
			playfieldArray[i-1][j] = playfieldArray[i][j]
		}
	}
}
