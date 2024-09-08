package users

import (
	"github.com/gin-gonic/gin"
	"github.com/neonshaman/bookstore_users-api/domain/users"
	"github.com/neonshaman/bookstore_users-api/services"
	"github.com/neonshaman/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

// CreateUser attempts to write a user to the persistence layer, encoding proper JSON response to context
func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// GetUser attempts to read user by user_id, encoding proper JSON response to context
func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement FindUser!")
}
