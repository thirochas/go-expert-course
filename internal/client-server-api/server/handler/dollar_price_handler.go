package handler

import (
	"encoding/json"
	"fmt"
	"go-expert-course/internal/client-server-api/server/client"
	"go-expert-course/internal/client-server-api/server/repository"
	"net/http"
)

type DollarPriceHandler struct {
	dollarPriceClient     client.IDDollarPriceClient
	dollarPriceRepository repository.IDollarPriceRepository
}

func NewDollarPriceHandler(dollarPriceClient client.IDDollarPriceClient, dollarPriceRepository repository.IDollarPriceRepository) *DollarPriceHandler {
	return &DollarPriceHandler{
		dollarPriceClient:     dollarPriceClient,
		dollarPriceRepository: dollarPriceRepository,
	}
}

func (d *DollarPriceHandler) Handle(w http.ResponseWriter, r *http.Request) {
	dollarPrice, err := d.dollarPriceClient.Get()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	select {
	case <-r.Context().Done():
		fmt.Println("Request canceled by client")
		http.Error(w, "request canceled by client", http.StatusRequestTimeout)
		return
	default:
	}

	err = d.dollarPriceRepository.Save(dollarPrice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(dollarPrice)
}
