package memory

import "WebApp/internal/core"

type UserRepository struct {
	users []*core.User
}

func NewUserRepository() *UserRepository {
	repository := &UserRepository{users: []*core.User{}}

	user1 := &core.User{
		ID:        0,
		FirstName: "Samat",
		LastName:  "Zaydullin",
	}

	user2 := &core.User{
		ID:        1,
		FirstName: "Azat",
		LastName:  "Zln",
	}

	repository.users = append(repository.users, user1)
	repository.users = append(repository.users, user2)

	return repository
}

func (repository *UserRepository) GetAll() []*core.User {
	return repository.users
}

func (repository *UserRepository) GetById(id int) *core.User {
	return repository.users[id]
}
