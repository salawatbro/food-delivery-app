package services

import "github.com/salawatbro/food-delivery-app/interfaces"

type NotificationService struct {
	repository interfaces.NotificationRepositoryInterface
}

func NewNotificationService(repository interfaces.NotificationRepositoryInterface) interfaces.NotificationServiceInterface {
	return &NotificationService{repository}
}
