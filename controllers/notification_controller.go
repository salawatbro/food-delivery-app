package controllers

import "github.com/salawatbro/food-delivery-app/interfaces"

type NotificationController struct {
	service interfaces.NotificationServiceInterface
}

func NewNotificationController(service interfaces.NotificationServiceInterface) interfaces.NotificationControllerInterface {
	return &NotificationController{service}
}
