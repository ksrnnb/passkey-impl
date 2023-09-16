package repository

import (
	"errors"

	"github.com/ksrnnb/passkey-impl/model"
)

type UserRepository struct {
	users []*model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: []*model.User{
			{
				Id:       "sample",
				Name:     "sample user",
				Password: "password",
			},
		},
	}
}

func (r *UserRepository) FindById(id string) (*model.User, error) {
	for _, u := range r.users {
		if u.Id == id {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}
