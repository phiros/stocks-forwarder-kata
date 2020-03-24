package prices

import (
	"github.com/phiros/go-http-averages-kata/internal/stock_development/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockUseCase struct {
	argumentRecorder []*domain.StockPrices
}

func NewMockUseCase() *mockUseCase {
	return &mockUseCase{}
}

func (u *mockUseCase) invocations() int {
	return len(u.argumentRecorder)
}

func (u *mockUseCase) CalcAndForwardAsPercentages(stockPrices *domain.StockPrices) error {
	u.argumentRecorder = append(u.argumentRecorder, stockPrices)
	return nil
}

func Test_HttpAdapterCanCallUseCase(t *testing.T) {
	useCaseMock := NewMockUseCase()
	var httpAdapter *HttpAdapter = NewHttpAdapter(useCaseMock)
	err := httpAdapter.
		calcAndForwardUseCase.
		CalcAndForwardAsPercentages(domain.NewStockPriceSequence("MSF"))
	assert.NoError(t, err)
	assert.NotNil(t, httpAdapter)
	assert.Equal(t, 1, useCaseMock.invocations())
}
