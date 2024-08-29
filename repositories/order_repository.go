package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/salawatbro/food-delivery-app/interfaces"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) interfaces.OrderRepositoryInterface {
	return &OrderRepository{db}
}
