package main

import (
	"github.com/phiros/go-http-averages-kata/internal/stock_development/pkg/adapters/in/prices"
	"github.com/phiros/go-http-averages-kata/internal/stock_development/pkg/adapters/out/percentages"
	"github.com/phiros/go-http-averages-kata/internal/stock_development/use_cases/calc_and_forward_percentage_change"
)

func main() {
	outPortToOtherService := percentages.NewHttpAPI("http://foo.invalid/changes")
	useCase := calc_and_forward_percentage_change.NewPercentageForwarder(outPortToOtherService)
	inPortHttpPricesApi := prices.NewHttpAdapter(useCase, "/prices")
	inPortHttpPricesApi.Run()
}
