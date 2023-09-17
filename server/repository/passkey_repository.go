package repository

import (
	"github.com/ksrnnb/passkey-impl/model"
)

type PasskeyRepository struct {
	passkeys []*model.Passkey
}

func NewPasskeyRepository() *PasskeyRepository {
	return &PasskeyRepository{
		passkeys: []*model.Passkey{},
	}
}

func (r PasskeyRepository) FindByUserId(userId string) (*model.Passkey, error) {
	for _, p := range r.passkeys {
		if p.UserId == userId {
			return p, nil
		}
	}
	return nil, ErrRecordNotFound
}
