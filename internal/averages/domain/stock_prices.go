package domain

type PricePerDay struct {
	Day   string
	Price int
}

type StockPriceSequence struct {
	Symbol      string
	StockPrices []*PricePerDay
}

func NewStockPriceSequence(s string) *StockPriceSequence {
	return &StockPriceSequence{
		Symbol:      s,
		StockPrices: []*PricePerDay{},
	}
}

func (s *StockPriceSequence) AddStockPrice(day string, price int) {
	s.StockPrices = append(s.StockPrices, &PricePerDay{
		Day:   day,
		Price: price,
	})
}
