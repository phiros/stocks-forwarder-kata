package calc_and_forward_percentage_change

import (
	"errors"
	"github.com/phiros/stocks-forwarder-kata/internal/stock_development/domain"
	"github.com/phiros/stocks-forwarder-kata/internal/stock_development/pkg/adapters/out/percentages"
	"github.com/shopspring/decimal"
)

type PercentageForwarder struct {
	percentageSink percentages.Port
}

func NewPercentageForwarder(percentageSink percentages.Port) *PercentageForwarder {
	return &PercentageForwarder{percentageSink: percentageSink}
}

func (f *PercentageForwarder) CalcAndForwardAsPercentages(stockPrices *domain.StockPrices) error {
	if len(stockPrices.StockPrices) < 2 {
		return errors.New("not enough stock price days")
	}
	symbol := stockPrices.Symbol
	prevDay := stockPrices.StockPrices[0]
	spc := domain.NewStockPercentageChanges(symbol)
	decimal.DivisionPrecision = 4
	for _, curDay := range stockPrices.StockPrices[1:] {
		day := curDay.Day
		change := decimal.NewFromInt(int64(curDay.Price - prevDay.Price)).
			Div(decimal.NewFromInt(int64(prevDay.Price))).
			Mul(decimal.NewFromInt(int64(100)))
		spc.AddPercentageChange(day, change)
		prevDay = curDay
	}
	return f.percentageSink.Send(spc)
}
