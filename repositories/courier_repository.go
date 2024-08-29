package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/salawatbro/food-delivery-app/interfaces"
)

type CourierRepository struct {
	db *sqlx.DB
}

func NewCourierRepository(db *sqlx.DB) interfaces.CourierRepositoryInterface {
	return &CourierRepository{db}
}
