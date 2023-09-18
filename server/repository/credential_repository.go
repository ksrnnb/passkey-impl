package repository

import (
	"github.com/ksrnnb/passkey-impl/model"
)

type CredentialRepository struct {
	creds []*model.Credential
}

func NewCredentialRepository() *CredentialRepository {
	return &CredentialRepository{
		creds: []*model.Credential{},
	}
}

func (r *CredentialRepository) FindByUserId(userId string) []*model.Credential {
	var creds []*model.Credential
	for _, c := range r.creds {
		if c.UserId == userId {
			creds = append(creds, c)
		}
	}
	return creds
}

// Add updates credential if it exists or creates credential if not exists
func (r *CredentialRepository) Add(cred *model.Credential) {
	for i, c := range r.creds {
		if c.Id == cred.Id {
			r.creds[i] = cred
			return
		}
	}
	r.creds = append(r.creds, cred)
}
