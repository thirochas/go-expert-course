package client

import (
	"context"
	"encoding/json"
	"fmt"
	"go-expert-course/internal/client-server-api/server/model"
	"io"
	"net/http"
	"time"
)

type IDDollarPriceClient interface {
	Get() (model.DollarPrice, error)
}

type DollarPriceClient struct {
	url string
}

func NewDollarPriceClient() IDDollarPriceClient {
	return &DollarPriceClient{
		url: "https://economia.awesomeapi.com.br/json/last/USD-BRL",
	}
}

func (d *DollarPriceClient) Get() (model.DollarPrice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", d.url, nil)
	if err != nil {
		return model.DollarPrice{}, fmt.Errorf("error creating request: %v", err)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return model.DollarPrice{}, fmt.Errorf("error getting dolar price: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.DollarPrice{}, fmt.Errorf("error reading response: %v", err)
	}

	var dollarPriceResponse model.DolarPriceResponse
	err = json.Unmarshal(body, &dollarPriceResponse)
	if err != nil {
		return model.DollarPrice{}, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return dollarPriceResponse.Data, nil
}
