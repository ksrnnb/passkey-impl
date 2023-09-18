package repository

import (
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/ksrnnb/passkey-impl/model"
)

type UserRepository struct {
	users []*model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: []*model.User{
			{
				Id:          "sample",
				Name:        "sample user",
				Password:    "password",
				Credentials: []webauthn.Credential{},
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
	return nil, ErrRecordNotFound
}

// Add updates user if it exists or creates user if not exists
func (r *UserRepository) Add(user *model.User) {
	for i, u := range r.users {
		if u.Id == user.Id {
			r.users[i] = user
			return
		}
	}
	r.users = append(r.users, user)
}
