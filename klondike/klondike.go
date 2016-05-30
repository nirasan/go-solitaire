package klondike

import (
	"math/rand"
	"time"
	"strings"
	"fmt"
	"math"
)

type Klondike struct {
	deck     Cards
	table    []Cards
	cursor   *Position
	selected *Position
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
		deck:      CreateDeck(),
		table: make([]Cards, cardsListMax, cardsListMax),
		cursor:    &Position{0, 0},
		selected:  nil,
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
	for i, v := range k.table {
		strs = append(strs, fmt.Sprintf("%d: %s", i, v.String()))
	}
	return strings.Join(strs, "\n")
}

func (k *Klondike) Init() {
	// waste
	k.table[waste] = Cards{}
	// foundation
	for i := 0; i < 4; i++ {
		n := i + foundation1
		k.table[n] = Cards{}
	}
	// column
	for i := 0; i < 7; i++ {
		n := i + column1
		cards := Cards{}
		for j := 0; j <= i; j++ {
			cards = append(cards, k.PickCard())
		}
		cards[i].Open = true
		k.table[n] = cards
	}
	// stock
	length := len(k.deck)
	for i := 0; i < length; i++ {
		k.table[stock] = append(k.table[stock], k.PickCard())
	}
	// cursor
	k.cursor.Row = 0
	k.cursor.Col = k.LastCol(stock)
}

func (k *Klondike) LastCol(i int) int {
	return int(math.Max(float64(len(k.table[i]) - 1), 0))
}

func (k *Klondike) PickCard() *Card {
	i := rand.Intn(len(k.deck))
	c := k.deck[i]
	k.deck = append(k.deck[:i], k.deck[i+1:]...)
	return c
}

func (k *Klondike) Select() {
	k.selected = &Position{k.cursor.Col, k.cursor.Row}
}

func (k *Klondike) CursorLeft() {
	row := k.cursor.Row - 1
	if row < 0 {
		row = column7
	}
	col := k.LastCol(row)
	k.cursor.Row, k.cursor.Col = row, col
}

func (k *Klondike) CursorRight() {
	row := k.cursor.Row + 1
	if row >= cardsListMax {
		row = stock
	}
	col := k.LastCol(row)
	k.cursor.Row, k.cursor.Col = row, col
}

