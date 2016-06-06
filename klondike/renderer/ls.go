package renderer

import (
	"fmt"
	"github.com/nirasan/go-solitaire/klondike"
	"github.com/nsf/termbox-go"
	"os/exec"
	"strings"
)

const (
	lsCardFormat = "/%s%2d"
	lsCardBack   = "/---"
	lsCardEmpty  = "/   "
	lsCardLength = 4
)

type LsRenderer struct {
	k         *klondike.Klondike
	colorFlag bool
	debugFlag bool
	err       error
	username  string
}

func NewLsRenderer(k *klondike.Klondike, c bool, d bool) *LsRenderer {
	username, err := exec.Command("whoami").Output()
	if err != nil {
		username = []byte("solitaire")
	}
	return &LsRenderer{k, c, d, nil, string(username)}
}

func (r *LsRenderer) SetError(e error) {
	r.err = e
}

func (r *LsRenderer) prefix_format(group string, length int) string {
	return fmt.Sprintf("drwxr-xr-x %s %-11s %2d ", r.username, group, length)
}

func (r *LsRenderer) pos(x, y int) (int, int) {
	offsetX, offsetY := len(r.prefix_format("", 1)), 0
	return offsetX + x*lsCardLength, offsetY + y
}

func (r *LsRenderer) Render() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	if r.k.IsClear() {
		r.RenderClear()
	} else {
		r.RenderGame()
	}
	if r.debugFlag {
		r.RenderDebug()
	}

	termbox.Flush()
}

func (r *LsRenderer) RenderGame() {
	// スコア
	drawStringDefault(0, 0, fmt.Sprintf("total %d", r.k.Score))

	// 場札
	r.renderPrefix(0, 1, "stock")
	r.renderLastCard(0, 0, 1)

	// 捨て札
	r.renderPrefix(1, 2, "waste")
	r.renderLastCard(1, 0, 2)

	// 組札
	r.renderPrefix(2, 3, "foundation1")
	r.renderLastCard(2, 0, 3)
	r.renderPrefix(3, 4, "foundation2")
	r.renderLastCard(3, 0, 4)
	r.renderPrefix(4, 5, "foundation3")
	r.renderLastCard(4, 0, 5)
	r.renderPrefix(5, 6, "foundation4")
	r.renderLastCard(5, 0, 6)

	// 場札
	r.renderPrefix(6, 7, "column1")
	r.renderColumn(6, 7)
	r.renderPrefix(7, 8, "column2")
	r.renderColumn(7, 8)
	r.renderPrefix(8, 9, "column3")
	r.renderColumn(8, 9)
	r.renderPrefix(9, 10, "column4")
	r.renderColumn(9, 10)
	r.renderPrefix(10, 11, "column5")
	r.renderColumn(10, 11)
	r.renderPrefix(11, 12, "column6")
	r.renderColumn(11, 12)
	r.renderPrefix(12, 13, "column7")
	r.renderColumn(12, 13)
}

func (r *LsRenderer) RenderClear() {
	drawStringDefault(0, 0, "GAME CLEAR")
	drawStringDefault(0, 1, fmt.Sprintf("SCORE: %d", r.k.Score))
	termbox.HideCursor()
}

func (r *LsRenderer) RenderDebug() {
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
}

func (r *LsRenderer) renderPrefix(row, y int, group string) {
	drawStringDefault(0, y, r.prefix_format(group, len(r.k.Table[row])))
}

func (r *LsRenderer) renderLastCard(row, x, y int) {
	r.renderCard(row, r.k.LastCol(row), x, y)
}

func (r *LsRenderer) renderColumn(row, y int) {
	if len(r.k.Table[row]) > 0 {
		for i := range r.k.Table[row] {
			r.renderCard(row, i, i, y)
		}
	} else {
		r.renderCard(row, 0, 0, y)
	}
}

func (r *LsRenderer) renderCard(row, col, x, y int) {
	str := lsCardEmpty
	color := colorBlack
	if len(r.k.Table[row]) > 0 {
		card := r.k.Table[row][col]
		if card.Open {
			str = fmt.Sprintf(lsCardFormat, card.Suit.String(), card.Num)
			if card.Suit == klondike.Hearts || card.Suit == klondike.Diamonds {
				color = colorRed
			}
		} else {
			str = lsCardBack
		}
	}
	x, y = r.pos(x, y)
	drawString(x, y, str, color, termbox.ColorDefault)
	if r.k.Cursor.Row == row && r.k.Cursor.Col == col {
		termbox.SetCursor(x, y)
	}
	if r.k.Selected != nil && r.k.Selected.Row == row && r.k.Selected.Col == col {
		highlightColor(x, y, lsCardLength)
	}
}
