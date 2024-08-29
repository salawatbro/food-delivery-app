package services

import "github.com/salawatbro/food-delivery-app/interfaces"

type ProductService struct {
	repository interfaces.ProductRepositoryInterface
}

func NewProductService(repository interfaces.ProductRepositoryInterface) interfaces.ProductServiceInterface {
	return &ProductService{repository}
}
