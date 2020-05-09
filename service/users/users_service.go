package users

import (
	"github.com/f4nt0md3v/bookstore_users-api/domain/users"
	"github.com/f4nt0md3v/bookstore_users-api/utils/date_utils"
	"github.com/f4nt0md3v/bookstore_users-api/utils/errors"
)

func CreateUser(u users.User) (*users.User, *errors.RestError) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	u.Status = users.StatusActive
	u.DateCreated = date_utils.GetNowDBFormat()
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

func UpdateUser(isPartial bool, u users.User) (*users.User, *errors.RestError) {
	current, err := GetUser(u.ID)
	if err != nil {
		return nil, err
	}

	// PATCH (partial update) or PUT (full update)
	if isPartial {
		// Update all non empty properties
		if u.FirstName != "" {
			current.FirstName = u.FirstName
		}
		if u.LastName != "" {
			current.LastName = u.LastName
		}
		if u.Email != "" {
			current.Email = u.Email
		}
	} else {
		current.FirstName = u.FirstName
		current.LastName = u.LastName
		current.Email = u.Email
	}

	if updErr := current.Update(); updErr != nil {
		return nil, updErr
	}
	return current, nil
}

func DeleteUser(uid int64) *errors.RestError {
	u := &users.User{ID: uid}
	return u.Delete()
}

func Search(status string) ([]users.User, *errors.RestError) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
