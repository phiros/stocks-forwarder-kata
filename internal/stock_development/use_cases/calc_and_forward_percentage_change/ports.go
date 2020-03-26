package calc_and_forward_percentage_change

import "github.com/phiros/stocks-forwarder-kata/internal/stock_development/domain"

type Port interface {
	CalcAndForwardAsPercentages(stockPrices *domain.StockPrices) error
}
