package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/f4nt0md3v/bookstore_users-api/domain/users"
	userservice "github.com/f4nt0md3v/bookstore_users-api/service/users"
	"github.com/f4nt0md3v/bookstore_users-api/utils/errors"
)

func getUserId(uidParam string) (int64, *errors.RestError) {
	uid, parseErr := strconv.ParseInt(uidParam, 10, 64)
	if parseErr != nil {
		return 0, errors.NewBadRequestError("user id must be a number")
	}
	return uid, nil
}

func Create(c *gin.Context) {
	var u users.User
	if err := c.ShouldBindJSON(&u); err != nil {
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

func Get(c *gin.Context) {
	uid, parseErr := getUserId(c.Param("id"))
	if parseErr != nil {
		c.JSON(parseErr.StatusCode, parseErr)
		return
	}
	u, getErr := userservice.GetUser(uid)
	if getErr != nil {
		c.JSON(getErr.StatusCode, getErr)
		return
	}
	c.JSON(http.StatusOK, u)
}

func Update(c *gin.Context) {
	uid, parseErr := getUserId(c.Param("id"))
	if parseErr != nil {
		c.JSON(parseErr.StatusCode, parseErr)
		return
	}
	var u users.User
	if err := c.ShouldBindJSON(&u); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	u.ID = uid

	// PATCH (partial update) or PUT (full update)
	isPartial := c.Request.Method == http.MethodPatch
	res, updErr := userservice.UpdateUser(isPartial, u)

	if updErr != nil {
		c.JSON(updErr.StatusCode, updErr)
		return
	}
	c.JSON(http.StatusOK, res)
}

func Delete(c *gin.Context) {
	uid, parseErr := getUserId(c.Param("id"))
	if parseErr != nil {
		c.JSON(parseErr.StatusCode, parseErr)
		return
	}
	if delErr := userservice.DeleteUser(uid); delErr != nil {
		c.JSON(delErr.StatusCode, delErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status":"deleted"})
}
