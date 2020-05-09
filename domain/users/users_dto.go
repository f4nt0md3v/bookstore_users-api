// Data transfer object
// Object transferring between controllers and DAO layer

package users

import (
	"strings"

	"github.com/f4nt0md3v/bookstore_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	ID             int64     `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	DateCreated    string    `json:"date_created"`
	Status         string    `json:"status"`
	Password       string    `json:"password"`
}

type Users []User

func (u *User) Validate() *errors.RestError {
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)

	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	u.Password = strings.TrimSpace(u.Password)
	if u.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}

	return nil
}
