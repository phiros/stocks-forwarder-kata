package percentages

import "github.com/phiros/stocks-forwarder-kata/internal/stock_development/domain"

type Port interface {
	Send(changes *domain.StockPercentageChanges) error
}
