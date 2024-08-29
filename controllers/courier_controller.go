package controllers

import "github.com/salawatbro/food-delivery-app/interfaces"

type CourierController struct {
	service interfaces.CourierServiceInterface
}

func NewCourierController(service interfaces.CourierServiceInterface) interfaces.CourierControllerInterface {
	return &CourierController{service}
}
