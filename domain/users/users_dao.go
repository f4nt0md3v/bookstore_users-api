// Data Access Object
// Data access layer working with persistence mechanism such as DB

package users

import (
	"fmt"

	"github.com/f4nt0md3v/bookstore_users-api/datasource/mysql/users_db"
	"github.com/f4nt0md3v/bookstore_users-api/utils/date_utils"
	"github.com/f4nt0md3v/bookstore_users-api/utils/errors"
	"github.com/f4nt0md3v/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = ?;"
	queryUpdateUser       = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?;"
	queryDeleteUser       = "DELETE FROM users WHERE id = ?;"
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
	if getErr := res.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
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
	insertResult, saveErr := stmt.Exec(u.FirstName, u.LastName, u.Email, u.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
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

func (u *User) Update() *errors.RestError {
	// Prepare and validate SQL statement before execution
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer func() {
		_ = stmt.Close()
	}()

	// Execute the SQL statement
	if _, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.ID); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (u *User) Delete() *errors.RestError {
	// Prepare and validate SQL statement before execution
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer func() {
		_ = stmt.Close()
	}()

	// Execute the SQL statement
	if _, err = stmt.Exec(u.ID); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}
