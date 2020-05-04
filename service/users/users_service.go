package users

import (
	"github.com/f4nt0md3v/bookstore_users-api/domain/users"
	"github.com/f4nt0md3v/bookstore_users-api/utils/errors"
)

func CreateUser(u users.User) (*users.User, *errors.RestError) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	if err := u.Save(); err != nil {
		return nil, err
	}

	return &u, nil
}

func GetUser(uid int64) (*users.User, *errors.RestError) {
	u := users.User{ID: uid}
	if err := u.Get(); err != nil {
		return nil, err
	}
	return &u, nil
}
