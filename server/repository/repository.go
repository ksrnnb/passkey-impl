package repository

import "errors"

const RepositoriesContextName = "Repositories"

var ErrRecordNotFound = errors.New("record not found")

var Repos Repositories

func init() {
	Repos = Repositories{
		UserRepository:       NewUserRepository(),
		CredentialRepository: NewCredentialRepository(),
	}
}

type Repositories struct {
	*UserRepository
	*CredentialRepository
}
