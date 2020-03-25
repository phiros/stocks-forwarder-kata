package percentages

import "github.com/phiros/go-http-averages-kata/internal/stock_development/domain"

type Port interface {
	Send(changes *domain.StockPercentageChanges) error
}
