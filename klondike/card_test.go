package klondike

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuit_IsDifferentColor(t *testing.T) {
	samples := []struct {
		S1   Suit
		S2   Suit
		Bool bool
	}{
		{Hearts, Hearts, false}, {Diamonds, Diamonds, false}, {Clubs, Clubs, false}, {Spades, Spades, false},
		{Hearts, Diamonds, false}, {Hearts, Clubs, true}, {Hearts, Spades, true},
		{Diamonds, Clubs, true}, {Diamonds, Spades, true},
		{Clubs, Spades, false},
	}
	for _, s := range samples {
		assert.Equal(t, s.Bool, s.S1.IsDifferentColor(s.S2), fmt.Sprintf("test failed: %v", s))
	}
}

func TestCard_CanPutInTableau(t *testing.T) {
	samples := []struct {
		Num1  uint8
		Suit1 Suit
		Num2  uint8
		Suit2 Suit
		Bool  bool
	}{
		{1, Hearts, 2, Clubs, true}, {12, Spades, 13, Diamonds, true}, {7, Spades, 8, Diamonds, true},
		{1, Spades, 2, Clubs, false}, {12, Hearts, 13, Diamonds, false},
		{13, Spades, 1, Diamonds, false}, {2, Hearts, 1, Clubs, false},
		{8, Spades, 1, Spades, false},
	}
	for _, s := range samples {
		c1, c2 := &Card{s.Num1, s.Suit1}, &Card{s.Num2, s.Suit2}
		assert.Equal(t, c1.CanPutInTableau(c2), s.Bool, fmt.Sprintf("test failed: %v", s))
	}
}

func TestCard_CanPutInFoundation(t *testing.T) {
	c := &Card{1, Hearts}
	assert.True(t, c.CanPutInFoundation(nil), "カードがなければAを設置できる")
	c.Num = 13
	assert.False(t, c.CanPutInFoundation(nil), "カードがなければA以外を設置できない")

	samples := []struct {
		Num1  uint8
		Suit1 Suit
		Num2  uint8
		Suit2 Suit
		Bool  bool
	}{
		{2, Hearts, 1, Hearts, true}, {13, Diamonds, 12, Diamonds, true},
		{1, Spades, 2, Spades, false}, {5, Clubs, 6, Clubs, false},
		{2, Diamonds, 1, Hearts, false}, {13, Clubs, 12, Spades, false},
	}
	for _, s := range samples {
		c1, c2 := &Card{s.Num1, s.Suit1}, &Card{s.Num2, s.Suit2}
		assert.Equal(t, c1.CanPutInFoundation(c2), s.Bool, fmt.Sprintf("test failed: %v", s))
	}
}
