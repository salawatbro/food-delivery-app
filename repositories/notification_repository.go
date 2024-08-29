package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/salawatbro/food-delivery-app/interfaces"
)

type NotificationRepository struct {
	db *sqlx.DB
}

func NewNotificationRepository(db *sqlx.DB) interfaces.NotificationRepositoryInterface {
	return &NotificationRepository{db}
}
