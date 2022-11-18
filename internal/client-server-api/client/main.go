package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type DollarPrice struct {
	Bid string `json:"bid"`
}

const cotacaoUrl = "http://localhost:8080/cotacao"

func main() {
	dollarPrice, err := getDollarPrice()
	if err != nil {
		panic(err)
	}

	f, err := os.Create("cotacao.txt")
	if err != nil {
		panic(fmt.Errorf("error creating file: %v", err))
	}
	//defer f.Close()

	length, err := f.Write([]byte(fmt.Sprintf("Dólar: %v", dollarPrice.Bid)))
	if err != nil {
		panic(fmt.Errorf("error writing to file: %v", err))
	}
	fmt.Printf("File cotacao.txt created successfully! Lenght: %d bytes\n", length)
}

func getDollarPrice() (DollarPrice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", cotacaoUrl, nil)
	if err != nil {
		return DollarPrice{}, fmt.Errorf("error creating request: %v", err)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return DollarPrice{}, fmt.Errorf("error getting dolar price: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return DollarPrice{}, fmt.Errorf("error reading response: %v", err)
	}

	var dollarPrice DollarPrice
	err = json.Unmarshal(body, &dollarPrice)
	if err != nil {
		return DollarPrice{}, fmt.Errorf("error unmarshalling response: %v", err)
	}
	return dollarPrice, nil
}
