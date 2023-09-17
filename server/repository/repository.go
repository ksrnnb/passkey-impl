package repository

import "errors"

const RepositoriesContextName = "Repositories"

var ErrRecordNotFound = errors.New("record not found")

type Repositories struct {
	*UserRepository
	*PasskeyRepository
}
