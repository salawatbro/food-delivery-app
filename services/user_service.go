package services

import "github.com/salawatbro/food-delivery-app/interfaces"

type UserService struct {
	repository interfaces.UserRepositoryInterface
}

func NewUserService(repository interfaces.UserRepositoryInterface) interfaces.UserServiceInterface {
	return &UserService{repository}
}
