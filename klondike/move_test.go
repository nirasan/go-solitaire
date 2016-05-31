package klondike

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestKlondike_Draw(t *testing.T) {
	k := NewKlondike()
	k.Init()

	assert.Equal(t, len(k.Table[stock]), 24)
	assert.Equal(t, len(k.Table[waste]), 0)
	err := k.Draw()
	assert.Nil(t, err)
	assert.Equal(t, len(k.Table[stock]), 23)
	assert.Equal(t, len(k.Table[waste]), 1)
	assert.Equal(t, k.Table[waste][0].Open, true)

	k.Table[stock] = k.Table[stock][:0]
	assert.Equal(t, len(k.Table[stock]), 0)
	err = k.Draw()
	assert.Equal(t, err, StockIsEmpty)
}

func TestKlondike_WasteToStock(t *testing.T) {
	k := NewKlondike()
	k.Init()

	err := k.WasteToStock()
	assert.Equal(t, err, StockIsNotEmpty)

	k.Draw()
	k.Table[stock] = Cards{}

	assert.Equal(t, len(k.Table[stock]), 0)
	assert.Equal(t, len(k.Table[waste]), 1)
	err = k.WasteToStock()
	assert.Nil(t, err)
	assert.Equal(t, len(k.Table[stock]), 1)
	assert.Equal(t, len(k.Table[waste]), 0)
	assert.Equal(t, k.Table[stock][0].Open, false)
}

func TestKlondike_MoveToFoundation(t *testing.T) {
	k := NewKlondike()
	k.Init()

	var err error

	err = k.MoveToFoundation(&Position{stock, 0}, &Position{foundation1, 0})
	assert.Equal(t, err, InvalidMovement)

	err = k.MoveToFoundation(&Position{waste, 0}, &Position{column1, 0})
	assert.Equal(t, err, InvalidMovement)

	err = k.MoveToFoundation(&Position{column2, 0}, &Position{foundation1, 0})
	assert.Equal(t, err, CardIsNotLastCol)

	k.Table[column2][1].Num = 2
	err = k.MoveToFoundation(&Position{column2, 1}, &Position{foundation1, 0})
	assert.Equal(t, err, CanNotPutInFoundation)

	k.Table[column2][1].Num = 1
	assert.Equal(t, len(k.Table[column2]), 2)
	assert.Equal(t, len(k.Table[foundation1]), 0)
	err = k.MoveToFoundation(&Position{column2, 1}, &Position{foundation1, 0})
	assert.Equal(t, len(k.Table[column2]), 1, "場札から組み札へ移動できた")
	assert.Equal(t, len(k.Table[foundation1]), 1, "場札から組み札へ移動できた")
	assert.Equal(t, k.Table[column2][0].Open, true, "移動後に一番上のカードをめくる")

	k.Draw()
	k.Table[waste][0].Num = 2
	k.Table[waste][0].Suit = k.Table[foundation1][0].Suit
	log.Println(k)

	err = k.MoveToFoundation(&Position{waste, k.LastCol(waste)}, &Position{foundation1, 0})
	assert.Nil(t, err)
	assert.Equal(t, len(k.Table[waste]), 0, "捨て札から組み札へ移動できた")
	assert.Equal(t, len(k.Table[foundation1]), 2, "捨て札から組み札へ移動できた")
}

func TestKlondike_MoveToColumn(t *testing.T) {
	k := NewKlondike()
	k.Init()

	var err error

	err = k.MoveToColumn(&Position{stock, 0}, &Position{column1, 0})
	assert.Equal(t, err, InvalidMovement)

	err = k.MoveToColumn(&Position{waste, 0}, &Position{foundation1, 0})
	assert.Equal(t, err, InvalidMovement)

	err = k.MoveToColumn(&Position{column2, 0}, &Position{column1, 0})
	assert.Equal(t, err, CardIsNotOpen)

	k.Draw()
	k.Table[waste][0].Num = k.Table[column1][0].Num
	err = k.MoveToColumn(&Position{waste, 0}, &Position{column1, 0})
	assert.Equal(t, err, CanNotPutInTableau)

	k.Table[waste][0].Num = k.Table[column1][0].Num + 1
	k.Table[waste][0].Suit = k.Table[column1][0].Suit
	err = k.MoveToColumn(&Position{waste, 0}, &Position{column1, 0})
	assert.Equal(t, err, CanNotPutInTableau)

	k.Table[waste][0].Suit = Hearts
	k.Table[column1][0].Suit = Clubs
	k.Table[waste][0].Num = 5
	k.Table[column1][0].Num = 6
	assert.Equal(t, len(k.Table[waste]), 1)
	assert.Equal(t, len(k.Table[column1]), 1)
	err = k.MoveToColumn(&Position{waste, 0}, &Position{column1, 0})
	assert.Nil(t, err, "捨て札から場札へ移動できた")
	assert.Equal(t, len(k.Table[waste]), 0)
	assert.Equal(t, len(k.Table[column1]), 2)

	k.Table[column2][1].Num = 7
	k.Table[column2][1].Suit = Diamonds
	err = k.MoveToColumn(&Position{column1, 0}, &Position{column2, 1})
	assert.Nil(t, err, "場札から場札へ複数枚移動できた")
	assert.Equal(t, len(k.Table[column1]), 0)
	assert.Equal(t, len(k.Table[column2]), 4)

	log.Println(k)
}

func TestKlondike_Move(t *testing.T) {
	k := NewKlondike()
	k.Init()

	var err error

	k.Cursor = &Position{0, 23}
	k.Selected = &Position{0, 23}
	err = k.Move()
	assert.Nil(t, err)
	assert.Equal(t, len(k.Table[waste]), 1, "1枚引けている")
}
