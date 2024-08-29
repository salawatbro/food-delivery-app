package controllers

import "github.com/salawatbro/food-delivery-app/interfaces"

type ProductController struct {
	service interfaces.ProductServiceInterface
}

func NewProductController(service interfaces.ProductServiceInterface) interfaces.ProductControllerInterface {
	return &ProductController{service}
}
