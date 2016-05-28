package klondike

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStockToWaste(t *testing.T) {
	s := &Stock{&Card{Num: 1}, &Card{Num: 2}, &Card{Num: 3}}
	w := &Waste{}

	var err error
	err = StockToWaste(s, w, 3)
	assert.NotNil(t, err)
	assert.Equal(t, err, InvalidArgs)

	err = StockToWaste(s, w, 2)
	assert.Nil(t, err)
	assert.Equal(t, len(*s), 2)
	assert.Equal(t, len(*w), 1)

	s = &Stock{&Card{Num: 1}, &Card{Num: 2}, &Card{Num: 3}}
	w = &Waste{}

	err = StockToWaste(s, w, 1)
	assert.Nil(t, err)
	assert.Equal(t, len(*s), 1)
	assert.Equal(t, len(*w), 2)
}