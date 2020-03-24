package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RepresentStockPriceSequence(t *testing.T) {
	var sps *StockPriceSequence = NewStockPriceSequence("MSF")
	assert.Equal(t, "MSF", sps.Symbol)
}
