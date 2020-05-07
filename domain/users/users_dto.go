// Data transfer object
// Object transferring between controllers and DAO layer

package users

import (
	"strings"

	"github.com/f4nt0md3v/bookstore_users-api/utils/errors"
)

type User struct {
	ID             int64     `json:"id"             binding:"-"`
	FirstName      string    `json:"first_name"     binding:"-"`
	LastName       string    `json:"last_name"      binding:"-"`
	Email          string    `json:"email"          binding:"-"`
	DateCreated    string    `json:"date_created"   binding:"-"`
}

type Users []User

func (u *User) Validate() *errors.RestError {
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
