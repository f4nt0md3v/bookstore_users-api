// Data transfer object
// Object transferring between controllers and DAO layer

package users

import (
	"strings"

	"github.com/f4nt0md3v/bookstore_users-api/utils/errors"
)

type User struct {
	ID             int64     `json:"id"             binding:"-"`
	FirstName      string    `json:"first_name"     binding:"required"`
	LastName       string    `json:"last_name"      binding:"required"`
	Email          string    `json:"email"          binding:"required"`
	DateCreated    string    `json:"date_created"   binding:"-"`
}

type Users []User

func (u *User) Validate() *errors.RestError {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
