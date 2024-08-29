package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/salawatbro/food-delivery-app/interfaces"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) interfaces.UserRepositoryInterface {
	return &UserRepository{db}
}
