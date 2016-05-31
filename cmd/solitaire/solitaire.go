package main

import (
	"github.com/nirasan/solitaire/klondike"
	"github.com/nsf/termbox-go"
	"strings"
	"flag"
)

var k *klondike.Klondike
var err error
var color = flag.Bool("color", true, "draw color charactor")

func main() {
	flag.Parse()

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
				drawCard(x, y, card.Suit.String(), card.NumString())
				//drawStringDefault(x, y, fmt.Sprintf("%s%2d", card.Suit.String(), card.Num))
			} else {
				drawStringDefault(x, y, "===")
			}
		}
	}
	// カーソル
	termbox.SetCursor(k.Cursor.Col*4, k.Cursor.Row)
	// ハイライト
	if k.Selected != nil {
		changeColor(k.Selected.Col, k.Selected.Row, 3, termbox.ColorBlack, termbox.ColorWhite)
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

func drawCard(x, y int, suit, num string) {
	fg := termbox.ColorDefault
	if *color && (suit == "H" || suit == "D") {
		fg = termbox.ColorRed
	}
	drawString(x, y, suit, fg, termbox.ColorDefault)
	drawString(x+1, y, num, termbox.ColorDefault, termbox.ColorDefault)
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
	start := width * y + x * 4
	for i := 0; i < length; i++ {
		cell := termbox.CellBuffer()[start+i]
		termbox.SetCell(x * 4 + i, y, cell.Ch, fg, bg)
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
