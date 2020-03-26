package main

import (
	"github.com/phiros/stocks-forwarder-kata/internal/stock_development/pkg/adapters/in/prices"
	"github.com/phiros/stocks-forwarder-kata/internal/stock_development/pkg/adapters/out/percentages"
	"github.com/phiros/stocks-forwarder-kata/internal/stock_development/use_cases/calc_and_forward_percentage_change"
	"os"
)

type Config struct {
	percentageApiUrl string
	pricesPath       string
}

func main() {
	config := Config{
		percentageApiUrl: "http://foo.invalid/changes",
		pricesPath:       "/prices",
	}
	pApi := os.Getenv("STOCKS_PERCENTAGE_API")
	if len(pApi) > 0 {
		config.percentageApiUrl = pApi
	}
	pricesPath := os.Getenv("STOCKS_PRICES_PATH")
	if len(pricesPath) > 0 {
		config.percentageApiUrl = pricesPath
	}

	outPortToOtherService := percentages.NewHttpAPI(config.percentageApiUrl)
	useCase := calc_and_forward_percentage_change.NewPercentageForwarder(outPortToOtherService)
	inPortHttpPricesApi := prices.NewHttpAdapter(useCase, config.pricesPath)
	inPortHttpPricesApi.Run()
}
