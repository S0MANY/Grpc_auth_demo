package repository

import (
	"grpcAuth/internal/domains/entities"
	"time"
)

type UserDatabase struct {
	db []entities.User
}

func NewUserDatabase() *UserDatabase {
	return &UserDatabase{
		db: []entities.User{
			{"123", "Andrew", "Andrew", time.Now()},
		},
	}
}

func (ud *UserDatabase) FindWithUsername(username string) bool {
	for _, v := range ud.db {
		if v.Username == username {
			return true
		}
	}
	return false
}

func (ud *UserDatabase) AddUser(u *entities.User) {
	ud.db = append(ud.db, *u)
}

func (ud *UserDatabase) CheckAccessWithPassword(username string, password string) bool {
	for _, v := range ud.db {
		if v.Username == username && v.Password == password {
			return true
		}
	}
	return false
}
