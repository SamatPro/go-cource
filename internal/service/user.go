package service

import (
	"WebApp/internal/core"
	"context"
)

type UserRepository interface {
	//GetAll() []*core.User
	GetById(ctx context.Context, id string) (*core.User, error)
}

type UserService struct {
	userRepository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{userRepository: repository}
}

//func (service *UserService) GetAll() []*core.User {
//	return service.userRepository.GetAll()
//}

func (service *UserService) GetById(ctx context.Context, id string) (*core.User, error) {
	return service.userRepository.GetById(ctx, id)
}
