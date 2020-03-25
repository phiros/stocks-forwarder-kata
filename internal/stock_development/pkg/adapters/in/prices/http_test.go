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

func Test_HttpAdapterPassesTransformedJsonDataToUseCase(t *testing.T) {
	useCaseMock := NewMockUseCase()
	var httpAdapter *HttpAdapter = NewHttpAdapter(useCaseMock)

	req, err := http.NewRequest("POST", "/prices", strings.NewReader(sampleGoodInputJson))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httpAdapter.httpHandler())

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.NotNil(t, httpAdapter)
	assert.Equal(t, 2, useCaseMock.invocations())
}

func Test_HttpAdapterErrorsOnJsonWithNoStocks(t *testing.T) {
	useCaseMock := NewMockUseCase()
	var httpAdapter *HttpAdapter = NewHttpAdapter(useCaseMock)

	emptyJSON := `{}`
	req, err := http.NewRequest("POST", "/prices", strings.NewReader(emptyJSON))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httpAdapter.httpHandler())

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.NotNil(t, httpAdapter)
	assert.Equal(t, 0, useCaseMock.invocations())
}

func Test_HttpAdapterErrorsOnJsonArray(t *testing.T) {
	useCaseMock := NewMockUseCase()
	var httpAdapter *HttpAdapter = NewHttpAdapter(useCaseMock)

	arrayJSON := `10`
	req, err := http.NewRequest("POST", "/prices", strings.NewReader(arrayJSON))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httpAdapter.httpHandler())

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.NotNil(t, httpAdapter)
	assert.Equal(t, 0, useCaseMock.invocations())
}

func Test_HttpAdapterErrorsOnJsonWithPriceChangeInWrongFormat(t *testing.T) {
	useCaseMock := NewMockUseCase()
	var httpAdapter *HttpAdapter = NewHttpAdapter(useCaseMock)

	jsonWithPriceChangeInWrongFormat := `{
"msf": [
[],
[]
]
}`
	req, err := http.NewRequest("POST", "/prices", strings.NewReader(jsonWithPriceChangeInWrongFormat))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httpAdapter.httpHandler())

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.NotNil(t, httpAdapter)
	assert.Equal(t, 0, useCaseMock.invocations())
}

func Test_HttpAdapterErrorsOnJsonWithPriceChangeNotParseable(t *testing.T) {
	useCaseMock := NewMockUseCase()
	var httpAdapter *HttpAdapter = NewHttpAdapter(useCaseMock)

	jsonWithPriceChangeInWrongFormat := `{
"msf": [
[120, 120],
[90, 90]
]
}`
	req, err := http.NewRequest("POST", "/prices", strings.NewReader(jsonWithPriceChangeInWrongFormat))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httpAdapter.httpHandler())

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.NotNil(t, httpAdapter)
	assert.Equal(t, 0, useCaseMock.invocations())
}

func Test_HttpAdapterErrorsOnJsonWithPriceChangeNotParseableCont(t *testing.T) {
	useCaseMock := NewMockUseCase()
	var httpAdapter *HttpAdapter = NewHttpAdapter(useCaseMock)

	jsonWithPriceChangeInWrongFormat := `{
"msf": [
["foo", "bar"],
["fizz", "buzz"]
]
}`
	req, err := http.NewRequest("POST", "/prices", strings.NewReader(jsonWithPriceChangeInWrongFormat))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httpAdapter.httpHandler())

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.NotNil(t, httpAdapter)
	assert.Equal(t, 0, useCaseMock.invocations())
}
