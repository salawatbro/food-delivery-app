package services

import "github.com/salawatbro/food-delivery-app/interfaces"

type OrderService struct {
	repository interfaces.OrderRepositoryInterface
}

func NewOrderService(repository interfaces.OrderRepositoryInterface) interfaces.OrderServiceInterface {
	return &OrderService{repository}
}
