package calc_and_forward_percentage_change

import (
	"github.com/phiros/go-http-averages-kata/internal/stock_development/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockPercentageSink struct {
}

func (mps *MockPercentageSink) Send(_ domain.StockPercentageChanges) {

}

func Test_CreateNewPercentageForwarder(t *testing.T) {
	mps := &MockPercentageSink{}
	var pf *PercentageForwarder = NewPercentageForwarder(mps)
	assert.Equal(t, pf.percentageSink, mps)
}
