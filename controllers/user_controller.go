package controllers

import "github.com/salawatbro/food-delivery-app/interfaces"

type UserController struct {
	service interfaces.UserServiceInterface
}

func NewUserController(service interfaces.UserServiceInterface) interfaces.UserControllerInterface {
	return &UserController{service}
}
