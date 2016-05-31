package klondike

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCreateDeck(t *testing.T) {
	deck := CreateDeck()
	assert.Equal(t, len(deck), 52)
}

func TestKlondike_PickCard(t *testing.T) {
	k := NewKlondike()
	c := k.PickCard()
	assert.NotNil(t, c)
	log.Println(c)
}

func TestKlondike_Init(t *testing.T) {
	k := NewKlondike()
	k.Init()
	assert.Equal(t, len(k.Table[stock]), 24)
	assert.Equal(t, len(k.Table[waste]), 0)
	assert.Equal(t, len(k.Table[foundation1]), 0)
	assert.Equal(t, len(k.Table[foundation2]), 0)
	assert.Equal(t, len(k.Table[foundation3]), 0)
	assert.Equal(t, len(k.Table[foundation4]), 0)
	assert.Equal(t, len(k.Table[column1]), 1)
	assert.Equal(t, len(k.Table[column2]), 2)
	assert.Equal(t, len(k.Table[column3]), 3)
	assert.Equal(t, len(k.Table[column4]), 4)
	assert.Equal(t, len(k.Table[column5]), 5)
	assert.Equal(t, len(k.Table[column6]), 6)
	assert.Equal(t, len(k.Table[column7]), 7)
	log.Println(k)
}

func TestKlondike_CursorLeft(t *testing.T) {
	k := NewKlondike()
	k.Init()
	samples := []struct {
		Row, Col int
	}{
		{0, 23}, {12, 6}, {11, 5}, {10, 4}, {9, 3}, {8, 2}, {7, 1}, {6, 0},
		{5, 0}, {4, 0}, {3, 0}, {2, 0}, {1, 0}, {0, 23}, {12, 6},
	}
	for _, s := range samples {
		assert.Equal(t, k.Cursor.Row, s.Row)
		assert.Equal(t, k.Cursor.Col, s.Col)
		k.CursorLeft()
	}
}

func TestKlondike_CursorRight(t *testing.T) {
	k := NewKlondike()
	k.Init()
	samples := []struct {
		Row, Col int
	}{
		{0, 23}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0},
		{6, 0}, {7, 1}, {8, 2}, {9, 3}, {10, 4}, {11, 5}, {12, 6}, {0, 23}, {1, 0},
	}
	for _, s := range samples {
		assert.Equal(t, k.Cursor.Row, s.Row)
		assert.Equal(t, k.Cursor.Col, s.Col)
		k.CursorRight()
	}
}

func TestKlondike_CursorUp(t *testing.T) {
	k := NewKlondike()
	k.Init()
	k.Table[column2][0].Open = true
	k.Table[column5][3].Open = true
	k.Table[column6][4].Open = true
	k.Table[column7][5].Open = true
	samples := []struct {
		Row, Col int
	}{
		{0, 23}, {12, 6}, {12, 5}, {11, 5}, {11, 4}, {10, 4}, {10, 3}, {9, 3}, {8, 2}, {7, 1}, {7, 0}, {6, 0},
		{5, 0}, {4, 0}, {3, 0}, {2, 0}, {1, 0}, {0, 23}, {12, 6},
	}
	for _, s := range samples {
		assert.Equal(t, k.Cursor.Row, s.Row)
		assert.Equal(t, k.Cursor.Col, s.Col)
		k.CursorUp()
	}
}

func TestKlondike_CursorDown(t *testing.T) {
	k := NewKlondike()
	k.Init()
	k.Table[column2][0].Open = true
	k.Table[column5][3].Open = true
	k.Table[column6][4].Open = true
	k.Table[column7][5].Open = true
	samples := []struct {
		Row, Col int
	}{
		{0, 23}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0},
		{6, 0}, {7, 0}, {7, 1}, {8, 2}, {9, 3}, {10, 3}, {10, 4}, {11, 4}, {11, 5}, {12, 5}, {12, 6}, {0, 23}, {1, 0},
	}
	for _, s := range samples {
		assert.Equal(t, k.Cursor.Row, s.Row)
		assert.Equal(t, k.Cursor.Col, s.Col)
		k.CursorDown()
	}
}
