package main

import (
	"fmt"
	"go-expert-course/internal/client-server-api/server/client"
	"go-expert-course/internal/client-server-api/server/handler"
	"go-expert-course/internal/client-server-api/server/model"
	"go-expert-course/internal/client-server-api/server/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

const dbFileName = "dollar_price.db"

func main() {
	db, err := gorm.Open(sqlite.Open(dbFileName), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("error opening database: %v", err))
	}

	db.AutoMigrate(&model.DollarPrice{})

	dollarPriceClient := client.NewDollarPriceClient()
	dollarPriceRepository := repository.NewDollarPriceRepository(db)

	dollarPriceHandler := handler.NewDollarPriceHandler(dollarPriceClient, dollarPriceRepository)

	http.HandleFunc("/cotacao", dollarPriceHandler.Handle)
	http.ListenAndServe(":8080", nil)
}
