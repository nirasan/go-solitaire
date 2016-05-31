package renderer

import (
	"fmt"
	"github.com/nirasan/solitaire/klondike"
	"github.com/nsf/termbox-go"
	"strings"
)

const (
	basicCardFormat = "[%s%2d]"
	basicCardBack   = "[---]"
	basicCardEmpty  = "[   ]"
	basicCardLength = 5
)

type BasicRenderer struct {
	k         *klondike.Klondike
	colorFlag bool
	err       error
}

func NewBasicRenderer(k *klondike.Klondike, c bool) *BasicRenderer {
	return &BasicRenderer{k, c, nil}
}

func (r *BasicRenderer) SetError(e error) {
	r.err = e
}

func (r *BasicRenderer) pos(x, y int) (int, int) {
	offsetX, offsetY := 1, 1
	return offsetX + x*(basicCardLength+1), offsetY + y*2
}

func (r *BasicRenderer) Render() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// 場札
	r.renderLastCard(0, 0, 0)

	// 捨て札
	r.renderLastCard(1, 1, 0)

	// 組札
	r.renderLastCard(2, 3, 0)
	r.renderLastCard(3, 4, 0)
	r.renderLastCard(4, 5, 0)
	r.renderLastCard(5, 6, 0)

	// 場札
	r.renderColumn(6)
	r.renderColumn(7)
	r.renderColumn(8)
	r.renderColumn(9)
	r.renderColumn(10)
	r.renderColumn(11)
	r.renderColumn(12)

	// エラー
	if r.err != nil {
		drawStringDefault(45, 0, r.err.Error())
		r.err = nil
	}
	// デバッグ
	debugStrings := strings.Split(r.k.String(), "\n")
	debugRow := 1
	for i, s := range debugStrings {
		drawStringDefault(45, debugRow+i, s)
	}

	termbox.Flush()
}

func (r *BasicRenderer) renderLastCard(row, x, y int) {
	r.renderCard(row, r.k.LastCol(row), x, y)
}

func (r *BasicRenderer) renderColumn(row int) {
	if len(r.k.Table[row]) > 0 {
		for i := range r.k.Table[row] {
			r.renderCard(row, i, row-6, i+2)
		}
	} else {
		r.renderCard(row, 0, row-6, 2)
	}
}

func (r *BasicRenderer) renderCard(row, col, x, y int) {
	str := basicCardEmpty
	color := colorBlack
	if len(r.k.Table[row]) > 0 {
		card := r.k.Table[row][col]
		if card.Open {
			str = fmt.Sprintf(basicCardFormat, card.Suit.String(), card.Num)
			if card.Suit == klondike.Hearts || card.Suit == klondike.Diamonds {
				color = colorRed
			}
		} else {
			str = basicCardBack
		}
	}
	x, y = r.pos(x, y)
	drawString(x, y, str, color, termbox.ColorDefault)
	if r.k.Cursor.Row == row && r.k.Cursor.Col == col {
		termbox.SetCursor(x, y)
	}
	if r.k.Selected != nil && r.k.Selected.Row == row && r.k.Selected.Col == col {
		highlightColor(x, y, basicCardLength)
	}
}
