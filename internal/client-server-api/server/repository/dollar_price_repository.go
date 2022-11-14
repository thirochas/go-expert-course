package repository

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"go-expert-course/internal/client-server-api/server/model"
	"gorm.io/gorm"
	"time"
)

type IDollarPriceRepository interface {
	Save(dollarPrice model.DollarPrice) error
}

type DollarPriceRepository struct {
	db *gorm.DB
}

func NewDollarPriceRepository(db *gorm.DB) IDollarPriceRepository {
	return &DollarPriceRepository{
		db: db,
	}
}

func (d *DollarPriceRepository) Save(dollarPrice model.DollarPrice) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	tx := d.db.WithContext(ctx).Create(&dollarPrice)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
