package prices

import (
	"encoding/json"
	"errors"
	"github.com/phiros/go-http-averages-kata/internal/stock_development/domain"
	"github.com/phiros/go-http-averages-kata/internal/stock_development/use_cases/calc_and_forward_percentage_change"
	"net/http"
)

type HttpAdapter struct {
	calcAndForwardUseCase calc_and_forward_percentage_change.Port
}

type MultipleStockPrices struct {
	Stocks []*domain.StockPrices
}

func NewMultipleStockPrices() *MultipleStockPrices {
	return &MultipleStockPrices{
		Stocks: []*domain.StockPrices{},
	}
}

func (m *MultipleStockPrices) UnmarshalJSON(data []byte) error {
	topLevel := map[string][][]interface{}{}
	if err := json.Unmarshal(data, &topLevel); err != nil {
		return err
	}

	if len(topLevel) == 0 {
		return errors.New("json in unexpected format")
	}

	for symbol, priceSequence := range topLevel {
		sp := domain.NewStockPriceSequence(symbol)
		for _, p := range priceSequence {
			if len(p) != 2 {
				return errors.New("json in unexpected format")
			}
			day, ok := p[0].(string)
			if !ok {
				return errors.New("json in unexpected format")
			}
			price, ok := p[1].(float64)
			if !ok {
				return errors.New("json in unexpected format")
			}
			sp.AddStockPrice(day, int(price))
		}
		m.Stocks = append(m.Stocks, sp)
	}
	return nil
}

func (a *HttpAdapter) httpHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		multipleStockPrices := NewMultipleStockPrices()
		err := json.NewDecoder(r.Body).Decode(&multipleStockPrices)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for _, sp := range multipleStockPrices.Stocks {
			_ = a.calcAndForwardUseCase.CalcAndForwardAsPercentages(sp)
		}
		w.WriteHeader(http.StatusOK)
	}
}

func NewHttpAdapter(calcAndForwardUseCase calc_and_forward_percentage_change.Port) *HttpAdapter {
	return &HttpAdapter{
		calcAndForwardUseCase: calcAndForwardUseCase,
	}
}
