package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
