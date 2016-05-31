package klondike

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

type Klondike struct {
	deck     Cards
	Table    []Cards
	Cursor   *Position
	Selected *Position
}

type Position struct {
	Row, Col int
}

const (
	stock = iota
	waste
	foundation1
	foundation2
	foundation3
	foundation4
	column1
	column2
	column3
	column4
	column5
	column6
	column7
	cardsListMax
)

func NewKlondike() *Klondike {
	rand.Seed(time.Now().UnixNano())
	return &Klondike{
		deck:     CreateDeck(),
		Table:    make([]Cards, cardsListMax, cardsListMax),
		Cursor:   &Position{0, 0},
		Selected: nil,
	}
}

func CreateDeck() Cards {
	cards := Cards{}
	for _, s := range []Suit{Hearts, Diamonds, Clubs, Spades} {
		for i := 1; i <= 13; i++ {
			cards = append(cards, NewCard(uint8(i), s))
		}
	}
	return cards
}

func (k *Klondike) String() string {
	var strs []string
	strs = append(strs, "CARDS:")
	for i, v := range k.Table {
		strs = append(strs, fmt.Sprintf("%d: %s", i, v.String()))
	}
	strs = append(strs, fmt.Sprintf("Cursor: %d, %d", k.Cursor.Col, k.Cursor.Row))
	if k.Selected == nil {
		strs = append(strs, "Selected: nil")
	} else {
		strs = append(strs, fmt.Sprintf("Selected: %d, %d", k.Selected.Col, k.Selected.Row))
	}
	return strings.Join(strs, "\n")
}

func (k *Klondike) Init() {
	// waste
	k.Table[waste] = Cards{}
	// foundation
	for i := 0; i < 4; i++ {
		n := i + foundation1
		k.Table[n] = Cards{}
	}
	// column
	for i := 0; i < 7; i++ {
		n := i + column1
		cards := Cards{}
		for j := 0; j <= i; j++ {
			cards = append(cards, k.PickCard())
		}
		cards[i].Open = true
		k.Table[n] = cards
	}
	// stock
	length := len(k.deck)
	for i := 0; i < length; i++ {
		k.Table[stock] = append(k.Table[stock], k.PickCard())
	}
	// cursor
	k.Cursor.Row = 0
	k.Cursor.Col = k.LastCol(stock)
}

func (k *Klondike) LastCol(i int) int {
	return int(math.Max(float64(len(k.Table[i])-1), 0))
}

func (k *Klondike) PickCard() *Card {
	i := rand.Intn(len(k.deck))
	c := k.deck[i]
	k.deck = append(k.deck[:i], k.deck[i+1:]...)
	return c
}

func (k *Klondike) Select() error {
	// 未選択
	if k.Selected == nil {
		k.Selected = &Position{Col: k.Cursor.Col, Row: k.Cursor.Row}
		return nil
	}
	// 選択済み
	err := k.Move()
	if err != nil {
		return err
	}
	return nil
}

func (k *Klondike) CursorLeft() {
	row := k.Cursor.Row - 1
	if row < 0 {
		row = column7
	}
	col := k.LastCol(row)
	k.Cursor.Row, k.Cursor.Col = row, col
}

func (k *Klondike) CursorRight() {
	row := k.Cursor.Row + 1
	if row >= cardsListMax {
		row = stock
	}
	col := k.LastCol(row)
	k.Cursor.Row, k.Cursor.Col = row, col
}

func (k *Klondike) CursorUp() {
	row := k.Cursor.Row
	col := k.Cursor.Col
	if stock <= row && row <= foundation4 {
		k.CursorLeft()
	} else {
		if col > 0 && k.Table[row][col-1].Open {
			k.Cursor.Col = col - 1
		} else {
			k.CursorLeft()
		}
	}
}

func (k *Klondike) CursorDown() {
	row := k.Cursor.Row
	col := k.Cursor.Col
	if stock <= row && row <= foundation4 {
		k.CursorRight()
	} else {
		if col < k.LastCol(row) && k.Table[row][col+1].Open {
			k.Cursor.Col = col + 1
		} else {
			if row == column7 {
				k.CursorRight()
			} else {
				for i, c := range k.Table[row+1] {
					if c.Open {
						k.Cursor.Row = row + 1
						k.Cursor.Col = i
						return
					}
				}
				k.CursorRight()
			}
		}
	}
}

func (k *Klondike) GetCard(p *Position) *Card {
	if len(k.Table[p.Row]) <= 0 {
		return nil
	}
	return k.Table[p.Row][p.Col]
}
