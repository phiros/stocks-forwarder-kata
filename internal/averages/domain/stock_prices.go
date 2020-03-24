package domain

type StockPriceSequence struct {
	Symbol string
}

func NewStockPriceSequence(s string) *StockPriceSequence {
	return &StockPriceSequence{Symbol: s}
}
