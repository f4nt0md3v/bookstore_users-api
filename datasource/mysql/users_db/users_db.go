package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUsersUsername = "MYSQL_USERNAME"
	mysqlUsersPassword = "MYSQL_PASSWORD"
	mysqlUsersHostname = "MYSQL_HOSTNAME"
	mysqlUsersScheme   = "MYSQL_DATABASE"
)

var (
	Client *sql.DB

	username = os.Getenv(mysqlUsersUsername)
	password = os.Getenv(mysqlUsersPassword)
	host     = os.Getenv(mysqlUsersHostname)
	scheme   = os.Getenv(mysqlUsersScheme)
)

func init() {
	ds := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		scheme,
	)

	var err error
	Client, err = sql.Open("mysql", ds)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("[INIT] database successfully configured")
}
