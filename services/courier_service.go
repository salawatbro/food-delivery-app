package services

import "github.com/salawatbro/food-delivery-app/interfaces"

type CourierService struct {
	repository interfaces.CourierRepositoryInterface
}

func NewCourierService(repository interfaces.CourierRepositoryInterface) interfaces.CourierServiceInterface {
	return &CourierService{repository}
}
