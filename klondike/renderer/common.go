package renderer

import "github.com/nsf/termbox-go"

func drawString(x, y int, s string, fg, bg termbox.Attribute) {
	for i, r := range s {
		termbox.SetCell(x+i, y, r, fg, bg)
	}
}

func drawStringDefault(x, y int, s string) {
	drawString(x, y, s, termbox.ColorDefault, termbox.ColorDefault)
}

func highlightColor(x, y, length int) {
	width, _ := termbox.Size()
	start := width*y + x*4
	for i := 0; i < length; i++ {
		cell := termbox.CellBuffer()[start+i]
		if cell.Fg == termbox.ColorDefault && cell.Bg == termbox.ColorDefault {
			termbox.SetCell(x*4+i, y, cell.Ch, termbox.ColorBlack, termbox.ColorWhite)
		} else {
			termbox.SetCell(x*4+i, y, cell.Ch, cell.Bg, cell.Fg)
		}
	}
}
