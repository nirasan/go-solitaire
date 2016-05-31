package renderer

import (
	"github.com/nirasan/solitaire/klondike"
	"github.com/nsf/termbox-go"
)

type BasicRenderer struct {
	k         *klondike.Klondike
	colorFlag bool
}

func NewBasicRenderer(k *klondike.Klondike, c bool) *BasicRenderer {
	return &BasicRenderer{k, c}
}

func (r *BasicRenderer) Render() {
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
				drawStringDefault(x, y, "===")
			}
		}
	}
	// カーソル
	termbox.SetCursor(r.k.Cursor.Col*4, r.k.Cursor.Row)
	// ハイライト
	if r.k.Selected != nil {
		highlightColor(r.k.Selected.Col, r.k.Selected.Row, 3)
	}
	// エラー
	if err != nil {
		drawStringDefault(0, len(r.k.Table), err.Error())
		err = nil
	}

	termbox.Flush()
}

func (r *BasicRenderer) drawCard(x, y int, suit, num string) {
	fg := colorBlack
	if r.colorFlag && (suit == "H" || suit == "D") {
		fg = colorRed
	}
	drawString(x, y, suit+num, fg, termbox.ColorDefault)
}
