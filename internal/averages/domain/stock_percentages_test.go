package domain

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RepresentStockPercentageChanges(t *testing.T) {
	var sps *StockPercentageChanges = NewStockPercentageChanges("MSF")
	assert.Equal(t, "MSF", sps.Symbol)
}

func Test_AddPercentageChanges(t *testing.T) {
	var sps *StockPercentageChanges = NewStockPercentageChanges("MSF")
	sps.AddPercentageChange("2020-02-21", decimal.NewFromFloat(10.0))
	firstPercentageChange := sps.PercentageChange[0]
	assert.Equal(t, "2020-02-21", firstPercentageChange.Day)
	assert.Equal(t, decimal.NewFromFloat(10.0), firstPercentageChange.ChangeInPercent)
}

func Test_AddMultiplePercentageChanges(t *testing.T) {
	var sps *StockPercentageChanges = NewStockPercentageChanges("MSF")
	sps.AddPercentageChange("2020-02-21", decimal.NewFromFloat(10.0))
	firstPercentageChange := sps.PercentageChange[0]
	assert.Equal(t, "2020-02-21", firstPercentageChange.Day)
	assert.Equal(t, decimal.NewFromFloat(10.0), firstPercentageChange.ChangeInPercent)
}

func Test_SupportFluidInterfaceForPercentageChanges(t *testing.T) {
	sps := NewStockPercentageChanges("MSF").
		AddPercentageChange("2020-02-21", decimal.NewFromFloat(10.0)).
		AddPercentageChange("2020-02-22", decimal.NewFromFloat(15.0))
	firstPercentageChange := sps.PercentageChange[0]
	secondPercentageChange := sps.PercentageChange[1]
	assert.Equal(t, "2020-02-21", firstPercentageChange.Day)
	assert.Equal(t, decimal.NewFromFloat(10.0), firstPercentageChange.ChangeInPercent)
	assert.Equal(t, "2020-02-22", secondPercentageChange.Day)
	assert.Equal(t, decimal.NewFromFloat(15.0), secondPercentageChange.ChangeInPercent)
}
