package users

import (
	"github.com/f4nt0md3v/bookstore_users-api/domain/users"
	"github.com/f4nt0md3v/bookstore_users-api/utils/crypto_utils"
	"github.com/f4nt0md3v/bookstore_users-api/utils/date_utils"
	"github.com/f4nt0md3v/bookstore_users-api/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestError)
	GetUser(int64) (*users.User, *errors.RestError)
	UpdateUser(bool, users.User) (*users.User, *errors.RestError)
	DeleteUser(int64) *errors.RestError
	SearchUser(string) (users.Users, *errors.RestError)
}

func (s *usersService) CreateUser(u users.User) (*users.User, *errors.RestError) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	u.Status = users.StatusActive
	u.DateCreated = date_utils.GetNowDBFormat()
	u.Password = crypto_utils.GetMD5(u.Password)
	if err := u.Save(); err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *usersService) GetUser(uid int64) (*users.User, *errors.RestError) {
	u := users.User{ID: uid}
	if err := u.Get(); err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *usersService) UpdateUser(isPartial bool, u users.User) (*users.User, *errors.RestError) {
	current, err := UsersService.GetUser(u.ID)
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

func (s *usersService) DeleteUser(uid int64) *errors.RestError {
	u := &users.User{ID: uid}
	return u.Delete()
}

func (s *usersService) SearchUser(status string) (users.Users, *errors.RestError) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
