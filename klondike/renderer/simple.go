package renderer

import (
	"github.com/nirasan/solitaire/klondike"
	"github.com/nsf/termbox-go"
	"strings"
)

type SimpleRenderer struct {
	k         *klondike.Klondike
	colorFlag bool
}

var (
	err        error
	colorRed   = termbox.Attribute(10)
	colorBlack = termbox.Attribute(243)
)

func NewSimpleRenderer(k *klondike.Klondike, c bool) *SimpleRenderer {
	return &SimpleRenderer{k, c}
}

func (r *SimpleRenderer) Render() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// カード
	for i, row := range r.k.Table {
		for j, card := range row {
			y := i
			x := j * 4
			if card == nil {
				continue
			}
			if card.Open {
				r.drawCard(x, y, card.Suit.String(), card.NumString())
			} else {
				r.drawStringDefault(x, y, "===")
			}
		}
	}
	// カーソル
	termbox.SetCursor(r.k.Cursor.Col*4, r.k.Cursor.Row)
	// ハイライト
	if r.k.Selected != nil {
		r.toggleColor(r.k.Selected.Col, r.k.Selected.Row, 3)
	}
	// エラー
	if err != nil {
		r.drawStringDefault(0, len(r.k.Table), err.Error())
		err = nil
	}
	// デバッグ
	debugStrings := strings.Split(r.k.String(), "\n")
	debugRow := len(r.k.Table) + 1
	for i, s := range debugStrings {
		r.drawStringDefault(0, debugRow+i, s)
	}

	termbox.Flush()
}

func (r *SimpleRenderer) drawCard(x, y int, suit, num string) {
	fg := colorBlack
	if r.colorFlag && (suit == "H" || suit == "D") {
		fg = colorRed
	}
	r.drawString(x, y, suit+num, fg, termbox.ColorDefault)
}

func (r *SimpleRenderer) drawString(x, y int, s string, fg, bg termbox.Attribute) {
	for i, r := range s {
		termbox.SetCell(x+i, y, r, fg, bg)
	}
}

func (r *SimpleRenderer) drawStringDefault(x, y int, s string) {
	r.drawString(x, y, s, termbox.ColorDefault, termbox.ColorDefault)
}

func (r *SimpleRenderer) toggleColor(x, y, length int) {
	width, _ := termbox.Size()
	start := width*y + x*4
	for i := 0; i < length; i++ {
		cell := termbox.CellBuffer()[start+i]
		termbox.SetCell(x*4+i, y, cell.Ch, cell.Bg, cell.Fg)
	}
}
