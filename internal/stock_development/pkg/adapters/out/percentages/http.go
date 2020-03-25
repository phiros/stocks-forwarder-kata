package percentages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/phiros/go-http-averages-kata/internal/stock_development/domain"
	"net/http"
)

type API struct {
	client  *http.Client
	baseUrl string
}

type Changes domain.StockPercentageChanges

func (changes *Changes) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("[")
	length := len(changes.PercentageChange)
	for i, v := range changes.PercentageChange {
		jsonDay, err := json.Marshal(v.Day)
		if err != nil {
			return nil, err
		}
		percentAsFloat, _ := v.ChangeInPercent.Float64()
		if err != nil {
			return nil, err
		}
		buffer.WriteString(fmt.Sprintf("[%s,%.2f]", jsonDay, percentAsFloat))
		if i < length-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("]")
	return buffer.Bytes(), nil
}

func (a *API) Send(changes *domain.StockPercentageChanges) error {
	c := Changes(*changes)
	b, err := json.Marshal(&c)
	if err != nil {
		return err
	}
	resp, err := a.client.Post(a.baseUrl+"/"+changes.Symbol, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	if resp.StatusCode >= 300 {
		return fmt.Errorf("percentages API returned: %d", resp.StatusCode)
	}
	return nil
}

func NewHttpAPIWithClient(baseUrl string, client *http.Client) *API {
	return &API{client, baseUrl}
}

func NewHttpAPI(baseUrl string) *API {
	return &API{http.DefaultClient, baseUrl}
}
