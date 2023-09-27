package service

import "WebApp/internal/core"

type UserRepository interface {
	GetAll() []*core.User
	GetById(id int) *core.User
}

type UserService struct {
	userRepository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{userRepository: repository}
}

func (service *UserService) GetAll() []*core.User {
	return service.userRepository.GetAll()
}

func (service *UserService) GetById(id int) *core.User {
	return service.userRepository.GetById(id)
}
