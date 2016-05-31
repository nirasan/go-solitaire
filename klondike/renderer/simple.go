package renderer

import (
	"github.com/nirasan/solitaire/klondike"
	"github.com/nsf/termbox-go"
	"strings"
)

type SimpleRenderer struct {
	k         *klondike.Klondike
	colorFlag bool
	err       error
}

const (
	colorRed   = termbox.Attribute(10)
	colorBlack = termbox.Attribute(243)
)

func NewSimpleRenderer(k *klondike.Klondike, c bool) *SimpleRenderer {
	return &SimpleRenderer{k, c, nil}
}

func (r *SimpleRenderer) SetError(e error) {
	r.err = e
}

func (r *SimpleRenderer) pos(x, y int) (int, int) {
	return x * 4, y
}

func (r *SimpleRenderer) Render() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// カード
	for i, row := range r.k.Table {
		for j, card := range row {
			x, y := r.pos(j, i)
			if card == nil {
				continue
			}
			if card.Open {
				r.drawCard(x, y, card.Suit.String(), card.NumString())
			} else {
				drawStringDefault(x, y, "===")
			}
		}
	}
	// カーソル
	termbox.SetCursor(r.k.Cursor.Col*4, r.k.Cursor.Row)
	// ハイライト
	if r.k.Selected != nil {
		x, y := r.pos(r.k.Selected.Col, r.k.Selected.Row)
		highlightColor(x, y, 3)
	}
	// エラー
	if r.err != nil {
		drawStringDefault(0, len(r.k.Table), r.err.Error())
		r.err = nil
	}
	// デバッグ
	debugStrings := strings.Split(r.k.String(), "\n")
	debugRow := len(r.k.Table) + 1
	for i, s := range debugStrings {
		drawStringDefault(0, debugRow+i, s)
	}

	termbox.Flush()
}

func (r *SimpleRenderer) drawCard(x, y int, suit, num string) {
	fg := colorBlack
	if r.colorFlag && (suit == "H" || suit == "D") {
		fg = colorRed
	}
	drawString(x, y, suit+num, fg, termbox.ColorDefault)
}
