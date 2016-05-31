package main

import (
	"fmt"
	"github.com/nirasan/solitaire/klondike"
	"github.com/nsf/termbox-go"
	"strings"
)

var k *klondike.Klondike
var err error

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	k = klondike.NewKlondike()
	k.Init()

	draw()

	pollEvent()
}

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// カード
	for i, row := range k.Table {
		for j, card := range row {
			y := i
			x := j * 4
			if card == nil {
				continue
			}
			if card.Open {
				drawStringDefault(x, y, fmt.Sprintf("%s%2d", card.Suit.String(), card.Num))
			} else {
				drawStringDefault(x, y, "===")
			}
		}
	}
	// カーソル
	termbox.SetCursor(k.Cursor.Col*4, k.Cursor.Row)
	// ハイライト
	if k.Selected != nil {
		changeColor(k.Selected.Col, k.Selected.Row, 3, termbox.ColorDefault, termbox.ColorWhite)
	}
	// エラー
	if err != nil {
		drawStringDefault(0, len(k.Table), err.Error())
		err = nil
	}
	// デバッグ
	debugStrings := strings.Split(k.String(), "\n")
	debugRow := len(k.Table) + 1
	for i, s := range debugStrings {
		drawStringDefault(0, debugRow+i, s)
	}

	termbox.Flush()
}

func drawString(x, y int, s string, fg, bg termbox.Attribute) {
	for i, r := range s {
		termbox.SetCell(x+i, y, r, fg, bg)
	}
}

func drawStringDefault(x, y int, s string) {
	drawString(x, y, s, termbox.ColorDefault, termbox.ColorDefault)
}

func changeColor(x, y, length int, fg, bg termbox.Attribute) {
	width, _ := termbox.Size()
	start := width*y + x*4
	for i := 0; i < length; i++ {
		cell := termbox.CellBuffer()[start+i]
		cell.Fg, cell.Bg = fg, bg
	}
}

func pollEvent() {
Loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break Loop
			case termbox.KeyEnter, termbox.KeySpace:
				err = k.Select()
			case termbox.KeyArrowUp:
				k.CursorUp()
			case termbox.KeyArrowDown:
				k.CursorDown()
			case termbox.KeyArrowLeft:
				k.CursorLeft()
			case termbox.KeyArrowRight:
				k.CursorRight()
			case termbox.KeyTab:
				k.CursorJump()
			}
			draw()
		}
	}
}
