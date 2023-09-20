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

func (r *CredentialRepository) FindById(credId string) (*model.Credential, error) {
	for _, c := range r.creds {
		if c.Id() == credId {
			return c, nil
		}
	}
	return nil, ErrRecordNotFound
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
		if c.Id() == cred.Id() {
			r.creds[i] = cred
			return
		}
	}
	r.creds = append(r.creds, cred)
}

// Delete deletes credential
func (r *CredentialRepository) Delete(credId string) {
	for i, c := range r.creds {
		if c.Id() == credId {
			// order is not important
			r.creds[i] = r.creds[len(r.creds)-1]
			r.creds = r.creds[:len(r.creds)-1]
			return
		}
	}
}
