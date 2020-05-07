package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"

	"github.com/f4nt0md3v/bookstore_users-api/utils/errors"
)

const (
	SQLDuplicateEntryCode = 1062
	SQLErrorNoRows        = "no rows in result set"
)

func ParseError(err error) *errors.RestError {
	// Smart MySQL error handling
	// Try to cast returned error to MySQLError type
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		// If cast fail => return generic error response
		if strings.Contains(err.Error(), SQLErrorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing db response")
	}
	// If cast success => return details error message
	switch sqlErr.Number {
		case SQLDuplicateEntryCode:
			return errors.NewBadRequestError("invalid data")
	}
	// If error number does not match => return generic error response
	return errors.NewInternalServerError("error processing request")
}
