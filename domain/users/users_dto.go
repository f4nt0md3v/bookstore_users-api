// Data transfer object
// Object transferring between controllers and DAO layer

package users

import (
	"strings"

	"github.com/f4nt0md3v/bookstore_users-api/utils/errors"
)

type User struct {
	ID             int64     `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	DateCreated    string    `json:"date_created"`
}

type Users []User

func (u *User) Validate() *errors.RestError {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
