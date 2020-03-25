package prices

import (
	"github.com/phiros/go-http-averages-kata/internal/stock_development/domain"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const sampleGoodInputJson = `{
"msf": [
["2020-02-20", 120],
["2020-02-21", 100],
["2020-02-22", 110],
["2020-02-23", 130]
],
"goog": [
["2020-02-20", 220],
["2020-02-21", 250],
["2020-02-22", 210],
["2020-02-23", 180]
]
}`

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

type TestCase struct {
	inputJson               string
	expectedStatusCode      int
	expectedMockInvocations int
}

func Test_JsonParsing(t *testing.T) {
	testCases := []TestCase{
		{sampleGoodInputJson, http.StatusOK, 2},
		{`{}`, http.StatusBadRequest, 0},
		{`10`, http.StatusBadRequest, 0},
		{`{"msf": [[],[]]}`, http.StatusBadRequest, 0},
		{`{"msf":[[120, 120],[90, 90]]}`, http.StatusBadRequest, 0},
		{`{"msf":[["foo", "bar"],["fizz", "buzz"]]}`, http.StatusBadRequest, 0},
	}

	for _, tc := range testCases {
		useCaseMock := NewMockUseCase()
		var httpAdapter *HttpAdapter = NewHttpAdapter(useCaseMock)

		req, err := http.NewRequest("POST", "/prices", strings.NewReader(tc.inputJson))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(httpAdapter.httpHandler())

		handler.ServeHTTP(rr, req)

		assert.Equal(t, tc.expectedStatusCode, rr.Code)
		assert.Equal(t, tc.expectedMockInvocations, useCaseMock.invocations())
	}
}
