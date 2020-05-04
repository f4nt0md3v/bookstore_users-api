// Data Access Object
// Data access layer working with persistence mechanism such as DB

package users

import (
	"fmt"

	"github.com/f4nt0md3v/bookstore_users-api/utils/date_utils"
	"github.com/f4nt0md3v/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (u *User) Get() *errors.RestError {
	res := usersDB[u.ID]
	if res == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", u.ID))
	}

	u.ID = res.ID
	u.FirstName = res.FirstName
	u.LastName = res.LastName
	u.Email = res.Email
	u.DateCreated = res.DateCreated

	return nil
}

func (u *User) Save() *errors.RestError {
	find := usersDB[u.ID]
	if find != nil {
		if find.Email == u.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", u.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", u.ID))
	}

	u.DateCreated = date_utils.GetNowString()
	usersDB[u.ID] = u

	return nil
}
