// Data Access Object
// Data access layer working with persistence mechanism such as DB

package users

import (
	"fmt"

	"github.com/f4nt0md3v/bookstore_users-api/datasource/mysql/users_db"
	"github.com/f4nt0md3v/bookstore_users-api/logger"
	"github.com/f4nt0md3v/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, status, password, date_created) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, status, date_created FROM users WHERE id = ?;"
	queryUpdateUser       = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id = ?;"
	queryDeleteUser       = "DELETE FROM users WHERE id = ?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?;"
)

func (u *User) Get() *errors.RestError {
	// Prepare and validate SQL statement before execution
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		// return errors.NewInternalServerError(err.Error())
		return errors.NewInternalServerError("database error")
	}
	defer func() {
		_ = stmt.Close()
	}()

	// Execute the SQL statement
	res := stmt.QueryRow(u.ID)
	if getErr := res.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Status, &u.DateCreated); getErr != nil {
		logger.Error("error when trying to get user by id", getErr)
		// return mysql_utils.ParseError(getErr)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (u *User) Save() *errors.RestError {
	// Prepare and validate SQL statement before execution
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying prepare insert user statement", err)
		// return errors.NewInternalServerError(err.Error())
		return errors.NewInternalServerError("database error")
	}
	defer func() {
		_ = stmt.Close()
	}()

	// Execute the SQL statement
	insertResult, saveErr := stmt.Exec(u.FirstName, u.LastName, u.Email, u.Status, u.Password, u.DateCreated)
	if saveErr != nil {
		logger.Error("error when trying to save user", saveErr)
		// return mysql_utils.ParseError(saveErr)
		return errors.NewInternalServerError("database error")
	}

	// Another way of executing statement
	// Above way is preferred since we validate query ahead of the execution
	// result, err := users_db.Client.Exec(queryInsertUser, u.FirstName, u.LastName, u.Email, u.DateCreated)

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating new user", err)
		// return errors.NewInternalServerError(
			// fmt.Sprintf("error when trying to save user: %s", err.Error()))
		return errors.NewInternalServerError("database error")
	}

	u.ID = userId
	return nil
}

func (u *User) Update() *errors.RestError {
	// Prepare and validate SQL statement before execution
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying prepare update user statement", err)
		// return errors.NewInternalServerError(err.Error())
		return errors.NewInternalServerError("database error")
	}
	defer func() {
		_ = stmt.Close()
	}()

	// Execute the SQL statement
	if _, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.ID); err != nil {
		logger.Error("error when trying to update user", err)
		// return mysql_utils.ParseError(err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (u *User) Delete() *errors.RestError {
	// Prepare and validate SQL statement before execution
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying prepare delete user statement", err)
		// return errors.NewInternalServerError(err.Error())
		return errors.NewInternalServerError("database error")
	}
	defer func() {
		_ = stmt.Close()
	}()

	// Execute the SQL statement
	if _, err = stmt.Exec(u.ID); err != nil {
		logger.Error("error when trying to delete user", err)
		// return mysql_utils.ParseError(err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (u *User) FindByStatus(status string) ([]User, *errors.RestError) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find users by status statement", err)
		// return nil, errors.NewInternalServerError(err.Error())
		return nil, errors.NewInternalServerError("database error")
	}
	defer func() {
		// Closing statement to make sure we won't run out of connections to DB
		_ = stmt.Close()
	}()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find users by status", err)
		// return nil, errors.NewInternalServerError(err.Error())
		return nil, errors.NewInternalServerError("database error")
	}
	defer func() {
		// Closing rows to make sure we won't run out of connections to DB
		_ = rows.Close()
	}()

	result := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status);
		   err != nil {
		   	logger.Error("error when trying to scan user row into user struct", err)
			// return nil, mysql_utils.ParseError(err)
			return nil, errors.NewInternalServerError("database error")
		}
		result = append(result, user)
	}
	if len(result) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return result, nil
}
