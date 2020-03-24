package prices

import "github.com/phiros/go-http-averages-kata/internal/stock_development/use_cases/calc_and_forward_percentage_change"

type HttpAdapter struct {
	calcAndForwardUseCase calc_and_forward_percentage_change.Port
}

func NewHttpAdapter(calcAndForwardUseCase calc_and_forward_percentage_change.Port) *HttpAdapter {
	return &HttpAdapter{
		calcAndForwardUseCase: calcAndForwardUseCase,
	}
}
