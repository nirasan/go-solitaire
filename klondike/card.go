package klondike

type Suit uint8

const (
	Hearts Suit = 1 << iota
	Diamonds
	Clubs
	Spades
)

type Card struct {
	Num  uint8
	Suit Suit
}

func (s Suit) IsDifferentColor(s2 Suit) bool {
	red := (Hearts | Diamonds)
	if s&red == 0 {
		return s2&red != 0
	} else {
		return s2&red == 0
	}
}

func (c *Card) CanPutInTableau(target *Card) bool {
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
