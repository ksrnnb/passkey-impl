package repository

import (
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
	var user *model.User
	for _, u := range r.users {
		if u.Id == id {
			user = u
			break
		}
	}
	if user == nil {
		return nil, ErrRecordNotFound
	}
	user.Credentials = Repos.CredentialRepository.FindByUserId(user.Id)
	return user, nil
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
