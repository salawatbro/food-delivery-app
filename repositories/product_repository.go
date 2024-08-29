package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/salawatbro/food-delivery-app/interfaces"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) interfaces.ProductRepositoryInterface {
	return &ProductRepository{db}
}
