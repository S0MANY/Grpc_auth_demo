package repository

import "grpcAuth/internal/domains/entities"

type UsersRepo interface {
	FindWithUsername(username string) bool
	AddUser(u *entities.User)
	CheckAccessWithPassword(username string, password string) bool
}

type Repository struct {
	Users UsersRepo
}

func NewReposiitory() *Repository {
	return &Repository{
		Users: NewUserDatabase(),
	}
}
