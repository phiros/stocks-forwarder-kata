package calc_and_forward_percentage_change

import (
	"github.com/phiros/go-http-averages-kata/internal/stock_development/domain"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockPercentageSink struct {
	argumentRecorder []*domain.StockPercentageChanges
}

func NewMockPercentageSink() *MockPercentageSink {
	return &MockPercentageSink{argumentRecorder: []*domain.StockPercentageChanges{}}
}

func (mps *MockPercentageSink) invocations() int {
	return len(mps.argumentRecorder)
}

func (mps *MockPercentageSink) Send(spc *domain.StockPercentageChanges) {
	mps.argumentRecorder = append(mps.argumentRecorder, spc)
}

func Test_CreateNewPercentageForwarder(t *testing.T) {
	mps := &MockPercentageSink{}
	var pf *PercentageForwarder = NewPercentageForwarder(mps)
	assert.Equal(t, pf.percentageSink, mps)
}

func Test_CalculateAndForwardChangeForTwoInputDays(t *testing.T) {
	mps := NewMockPercentageSink()
	var pf *PercentageForwarder = NewPercentageForwarder(mps)
	prices := domain.NewStockPriceSequence("MSF").
		AddStockPrice("2020-02-20", 100).
		AddStockPrice("2020-02-21", 110)
	err := pf.CalcAndForwardAsPercentages(prices)

	assert.NoError(t, err)
	assert.Equal(t, mps.invocations(), 1)
	changes := mps.argumentRecorder[0].PercentageChange
	assert.Equal(t, changes[0].Day, "2020-02-21")
	assert.Equal(t, changes[0].ChangeInPercent, decimal.NewFromFloat(10))
}
