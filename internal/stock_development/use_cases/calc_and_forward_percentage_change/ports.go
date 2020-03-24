package calc_and_forward_percentage_change

import "github.com/phiros/go-http-averages-kata/internal/stock_development/domain"

type Port interface {
	CalcAndForwardAsPercentages(stockPrices *domain.StockPrices) error
}
