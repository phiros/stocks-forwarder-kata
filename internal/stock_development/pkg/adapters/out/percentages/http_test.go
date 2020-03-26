package percentages

import (
	"github.com/phiros/stocks-forwarder-kata/internal/stock_development/domain"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

// Note: the round tripper design for testing http clients
// is inspired by
// http://hassansin.github.io/Unit-Testing-http-client-in-Go
// I really recommend reading the article if you want to
// how this works

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func Test_SuccessfulSendWithoutJsonSchemaValidation(t *testing.T) {
	mockClientInvocations := 0
	client := NewTestClient(func(req *http.Request) *http.Response {
		mockClientInvocations++
		assert.Equal(t, "http://foo.invalid/changes/MSF", req.URL.String())
		assert.Equal(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
		}
	})

	api := NewHttpAPIWithClient("http://foo.invalid/changes", client)
	err := api.Send(domain.
		NewStockPercentageChanges("MSF").
		AddPercentageChange("2020-02-21", decimal.NewFromInt(10)))
	assert.NoError(t, err)
	assert.Equal(t, mockClientInvocations, 1)
}

func Test_SuccessfulSendWithJsonSchemaValidation(t *testing.T) {
	goodJson := `[["2020-02-21",-16.67],["2020-02-22",10.00],["2020-02-23",18.18]]`
	mockClientInvocations := 0
	client := NewTestClient(func(req *http.Request) *http.Response {
		mockClientInvocations++
		body, _ := ioutil.ReadAll(req.Body)
		assert.Equal(t, strings.TrimSpace(goodJson), strings.TrimSpace(string(body)))
		assert.Equal(t, "http://foo.invalid/changes/MSF", req.URL.String())
		assert.Equal(t, "POST", req.Method)
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
		}
	})

	api := NewHttpAPIWithClient("http://foo.invalid/changes", client)
	err := api.Send(domain.
		NewStockPercentageChanges("MSF").
		AddPercentageChange("2020-02-21", decimal.NewFromFloat(-16.67)).
		AddPercentageChange("2020-02-22", decimal.NewFromFloat(10)).
		AddPercentageChange("2020-02-23", decimal.NewFromFloat(18.18)))
	assert.NoError(t, err)
	assert.Equal(t, mockClientInvocations, 1)
}
