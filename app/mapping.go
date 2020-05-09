package app

import (
	"github.com/f4nt0md3v/bookstore_users-api/controller/ping"
	"github.com/f4nt0md3v/bookstore_users-api/controller/users"
)

func mapUrls()  {
	r.GET("/ping", ping.Ping)

	r.POST("/users", users.Create)
	r.GET("/users/:id", users.Get)
	r.PUT("/users/:id", users.Update)
	r.PATCH("/users/:id", users.Update)
	r.DELETE("/users/:id", users.Delete)
}
