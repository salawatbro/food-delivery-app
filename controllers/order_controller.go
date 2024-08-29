package controllers

import "github.com/salawatbro/food-delivery-app/interfaces"

type OrderController struct {
	service interfaces.OrderServiceInterface
}

func NewOrderController(service interfaces.OrderServiceInterface) interfaces.OrderControllerInterface {
	return &OrderController{service}
}
