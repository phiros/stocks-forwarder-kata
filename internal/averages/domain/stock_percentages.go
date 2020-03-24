package domain

import "github.com/shopspring/decimal"

type ChangePerDay struct {
	Day             string
	ChangeInPercent decimal.Decimal
}

type StockPercentageChanges struct {
	Symbol           string
	PercentageChange []*ChangePerDay
}

func (s *StockPercentageChanges) AddPercentageChange(day string, change decimal.Decimal) *StockPercentageChanges {
	s.PercentageChange = append(s.PercentageChange, &ChangePerDay{
		Day:             day,
		ChangeInPercent: change,
	})
	return s
}

func NewStockPercentageChanges(symbol string) *StockPercentageChanges {
	return &StockPercentageChanges{Symbol: symbol, PercentageChange: []*ChangePerDay{}}
}
