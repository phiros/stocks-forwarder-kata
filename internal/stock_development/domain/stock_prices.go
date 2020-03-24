package domain

type PricePerDay struct {
	Day   string
	Price int
}

type StockPrices struct {
	Symbol      string
	StockPrices []*PricePerDay
}

func NewStockPriceSequence(s string) *StockPrices {
	return &StockPrices{
		Symbol:      s,
		StockPrices: []*PricePerDay{},
	}
}

func (s *StockPrices) AddStockPrice(day string, price int) *StockPrices {
	s.StockPrices = append(s.StockPrices, &PricePerDay{
		Day:   day,
		Price: price,
	})
	return s
}
