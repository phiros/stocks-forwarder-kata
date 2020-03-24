package calc_and_forward_percentage_change

import "github.com/phiros/go-http-averages-kata/internal/stock_development/pkg/adapters/out/percentage_sink"

type PercentageForwarder struct {
	percentageSink percentage_sink.Port
}

func NewPercentageForwarder(percentageSink percentage_sink.Port) *PercentageForwarder {
	return &PercentageForwarder{percentageSink: percentageSink}
}
