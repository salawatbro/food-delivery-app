package services

import "github.com/salawatbro/food-delivery-app/interfaces"

type AuthService struct {
	repository interfaces.UserRepositoryInterface
}

func NewAuthService(repository interfaces.UserRepositoryInterface) interfaces.AuthServiceInterface {
	return &AuthService{repository}
}
