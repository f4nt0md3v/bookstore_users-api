package users

import (
	"github.com/f4nt0md3v/bookstore_users-api/domain/users"
	"github.com/f4nt0md3v/bookstore_users-api/utils/errors"
)

func CreateUser(u users.User) (*users.User, *errors.RestError) {
	return &u, nil
}
