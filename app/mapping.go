package app

import (
	"github.com/f4nt0md3v/bookstore_users-api/controller/ping"
	"github.com/f4nt0md3v/bookstore_users-api/controller/users"
)

func mapUrls()  {
	r.GET("/ping", ping.Ping)

	r.GET("/users/:id", users.GetUser)
	r.POST("/users", users.CreateUser)
}
