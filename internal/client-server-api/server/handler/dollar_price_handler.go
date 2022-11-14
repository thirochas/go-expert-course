package handler

import (
	"encoding/json"
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

func (d *DollarPriceHandler) Handle(w http.ResponseWriter, _ *http.Request) {
	dollarPrice, err := d.dollarPriceClient.Get()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d.dollarPriceRepository.Save(dollarPrice)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(dollarPrice)
}
