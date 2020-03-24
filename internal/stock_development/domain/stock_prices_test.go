package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RepresentStockPriceSequence(t *testing.T) {
	var sps *StockPriceSequence = NewStockPriceSequence("MSF")
	assert.Equal(t, "MSF", sps.Symbol)
}

func Test_AddStockPriceSequence(t *testing.T) {
	var sps *StockPriceSequence = NewStockPriceSequence("MSF")
	sps.AddStockPrice("2020-02-20", 100)
	assert.Equal(t, "2020-02-20", sps.StockPrices[0].Day)
	assert.Equal(t, 100, sps.StockPrices[0].Price)
}

func Test_AddMultipleStockPrices(t *testing.T) {
	var sps *StockPriceSequence = NewStockPriceSequence("MSF")
	sps.AddStockPrice("2020-02-20", 100)
	sps.AddStockPrice("2020-02-21", 110)
	assert.Equal(t, "2020-02-20", sps.StockPrices[0].Day)
	assert.Equal(t, 100, sps.StockPrices[0].Price)
	assert.Equal(t, "2020-02-21", sps.StockPrices[1].Day)
	assert.Equal(t, 110, sps.StockPrices[1].Price)
}

func Test_SupportFluidInterface(t *testing.T) {
	sps := NewStockPriceSequence("MSF").
		AddStockPrice("2020-02-20", 100).
		AddStockPrice("2020-02-21", 110)
	assert.Equal(t, "2020-02-20", sps.StockPrices[0].Day)
	assert.Equal(t, 100, sps.StockPrices[0].Price)
	assert.Equal(t, "2020-02-21", sps.StockPrices[1].Day)
	assert.Equal(t, 110, sps.StockPrices[1].Price)
}
