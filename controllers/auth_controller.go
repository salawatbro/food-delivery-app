package controllers

import "github.com/salawatbro/food-delivery-app/interfaces"

type AuthController struct {
	service interfaces.AuthServiceInterface
}

func NewAuthController(service interfaces.AuthServiceInterface) interfaces.AuthControllerInterface {
	return &AuthController{service}
}
