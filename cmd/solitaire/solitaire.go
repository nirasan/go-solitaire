package main

import (
	"flag"
	keyboard "github.com/jteeuwen/keyboard/termbox"
	"github.com/nirasan/solitaire/klondike"
	"github.com/nsf/termbox-go"
	"strings"
)

var (
	k          *klondike.Klondike
	err        error
	colorFlag  = flag.Bool("color", true, "draw color charactor")
	colorRed   = termbox.Attribute(10)
	colorBlack = termbox.Attribute(243)
)

func main() {
	flag.Parse()

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetOutputMode(termbox.Output256)

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
			} else {
				drawStringDefault(x, y, "===")
			}
		}
	}
	// カーソル
	termbox.SetCursor(k.Cursor.Col*4, k.Cursor.Row)
	// ハイライト
	if k.Selected != nil {
		toggleColor(k.Selected.Col, k.Selected.Row, 3)
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
	fg := colorBlack
	if *colorFlag && (suit == "H" || suit == "D") {
		fg = colorRed
	}
	drawString(x, y, suit+num, fg, termbox.ColorDefault)
}

func drawString(x, y int, s string, fg, bg termbox.Attribute) {
	for i, r := range s {
		termbox.SetCell(x+i, y, r, fg, bg)
	}
}

func drawStringDefault(x, y int, s string) {
	drawString(x, y, s, termbox.ColorDefault, termbox.ColorDefault)
}

func toggleColor(x, y, length int) {
	width, _ := termbox.Size()
	start := width*y + x*4
	for i := 0; i < length; i++ {
		cell := termbox.CellBuffer()[start+i]
		termbox.SetCell(x*4+i, y, cell.Ch, cell.Bg, cell.Fg)
	}
}

func pollEvent() {
	running := true

	kb := keyboard.New()
	kb.Bind(func() { running = false }, "escape", "q")
	kb.Bind(func() { k.CursorUp(); draw() }, "up", "k")
	kb.Bind(func() { k.CursorDown(); draw() }, "down", "j")
	kb.Bind(func() { k.CursorLeft(); draw() }, "left", "h")
	kb.Bind(func() { k.CursorRight(); draw() }, "right", "l")
	kb.Bind(func() { k.CursorJump(); draw() }, "tab")
	kb.Bind(func() { k.Select(); draw() }, "space")

	for running {
		kb.Poll(termbox.PollEvent())
	}
}
