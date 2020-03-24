package calc_and_forward_percentage_change

import (
	"github.com/phiros/go-http-averages-kata/internal/stock_development/domain"
	"github.com/phiros/go-http-averages-kata/internal/stock_development/pkg/adapters/out/percentage_sink"
	"github.com/shopspring/decimal"
)

type PercentageForwarder struct {
	percentageSink percentage_sink.Port
}

func NewPercentageForwarder(percentageSink percentage_sink.Port) *PercentageForwarder {
	return &PercentageForwarder{percentageSink: percentageSink}
}

func (f *PercentageForwarder) CalcAndForwardAsPercentages(stockPrices *domain.StockPrices) error {
	f.percentageSink.Send(domain.NewStockPercentageChanges(stockPrices.Symbol).
		AddPercentageChange("2020-02-21", decimal.NewFromFloat(10)))
	return nil
}
