package app

import (
	"github.com/f4nt0md3v/bookstore_users-api/controller/ping"
	"github.com/f4nt0md3v/bookstore_users-api/controller/users"
)

func mapUrls()  {
	r.GET("/ping", ping.Ping)

	r.POST("/users", users.CreateUser)
	r.GET("/users/:id", users.GetUser)
	r.PUT("/users/:id", users.UpdateUser)
	r.PATCH("/users/:id", users.UpdateUser)
	r.DELETE("/users/:id", users.DeleteUser)
}
