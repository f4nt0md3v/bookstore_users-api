package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/f4nt0md3v/bookstore_users-api/domain/users"
	userservice "github.com/f4nt0md3v/bookstore_users-api/service/users"
	"github.com/f4nt0md3v/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var u users.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}
	res, saveErr := userservice.CreateUser(u)
	if saveErr != nil {
		c.JSON(saveErr.StatusCode, saveErr)
		return
	}
	c.JSON(http.StatusCreated, res)
}

func GetUser(c *gin.Context) {
	uid, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		err := errors.NewBadRequestError("user id must be a number")
		c.JSON(err.StatusCode, err)
		return
	}
	u, getErr := userservice.GetUser(uid)
	if getErr != nil {
		c.JSON(getErr.StatusCode, getErr)
		return
	}
	c.JSON(http.StatusOK, u)
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
