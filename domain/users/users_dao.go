// Data Access Object
// Data access layer working with persistence mechanism such as DB

package users

import (
	"fmt"
	"strings"

	"github.com/f4nt0md3v/bookstore_users-api/datasource/mysql/users_db"
	"github.com/f4nt0md3v/bookstore_users-api/utils/date_utils"
	"github.com/f4nt0md3v/bookstore_users-api/utils/errors"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = ?;"
)

func (u *User) Get() *errors.RestError {
	// Prepare and validate SQL statement before execution
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer func() {
		_ = stmt.Close()
	}()

	// Execute the SQL statement
	res := stmt.QueryRow(u.ID)
	if err := res.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(
				fmt.Sprintf("user %d not found", u.ID))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user with id %d: %s", u.ID, err.Error()))
	}

	return nil
}

func (u *User) Save() *errors.RestError {
	// Prepare and validate SQL statement before execution
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer func() {
		_ = stmt.Close()
	}()

	u.DateCreated = date_utils.GetNowString()

	// Execute the SQL statement
	insertResult, err := stmt.Exec(u.FirstName, u.LastName, u.Email, u.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(
				fmt.Sprintf("email %s already exists", u.Email))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	// Another way of executing statement
	// Above way is preferred since we validate query ahead of the execution
	// result, err := users_db.Client.Exec(queryInsertUser, u.FirstName, u.LastName, u.Email, u.DateCreated)

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	u.ID = userId
	return nil
}
