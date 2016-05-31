package klondike

import (
	"fmt"
	"strconv"
	"strings"
)

type Suit uint8

const (
	Hearts Suit = 1 << iota
	Diamonds
	Clubs
	Spades
)

var suitLabel = map[Suit]string{
	Hearts:   "H",
	Diamonds: "D",
	Clubs:    "C",
	Spades:   "S",
}

func (s Suit) String() string {
	return suitLabel[s]
}

func (s Suit) IsDifferentColor(s2 Suit) bool {
	red := (Hearts | Diamonds)
	if s&red == 0 {
		return s2&red != 0
	} else {
		return s2&red == 0
	}
}

type Card struct {
	Num  uint8
	Suit Suit
	Open bool
}

var numberLabel = map[uint8]string{
	1: "A", 11: "J", 12: "Q", 13: "K",
}

func NewCard(n uint8, s Suit) *Card {
	return &Card{Num: n, Suit: s, Open: false}
}

func (c *Card) String() string {
	f := "(%s:%s)"
	if c.Open {
		f = "[%s:%s]"
	}
	return fmt.Sprintf(f, c.Suit.String(), c.NumString())
}

func (c *Card) NumString() string {
	s := strconv.Itoa(int(c.Num))
	if v, ok := numberLabel[c.Num]; ok {
		s = v
	}
	return s
}

func (c *Card) CanPutInTableau(target *Card) bool {
	if target == nil {
		return c.Num == 13
	}
	if !c.Suit.IsDifferentColor(target.Suit) {
		return false
	}
	if c.Num+1 != target.Num {
		return false
	}
	return true
}

func (c *Card) CanPutInFoundation(target *Card) bool {
	if target == nil {
		return c.Num == 1
	}
	if c.Suit != target.Suit {
		return false
	}
	if c.Num-1 != target.Num {
		return false
	}
	return true
}

type Cards []*Card

func (c Cards) String() string {
	var strs []string
	for _, v := range c {
		strs = append(strs, v.String())
	}
	return strings.Join(strs, " ")
}
