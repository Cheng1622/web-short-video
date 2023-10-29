package app

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const ContextEmailKey = "Email"
const ContextUserIdKey = "UserId"

var (
	ErrorNotLogin = errors.New("no login")
)

func GetCurrentUserId(c *gin.Context) (int64, error) {
	userId, ok := c.Get(ContextUserIdKey)
	if !ok {
		return 0, ErrorNotLogin
	}
	return userId.(int64), nil
}

func GetCurrentUserName(c *gin.Context) (string, error) {
	userName, ok := c.Get(ContextUserIdKey)
	if !ok {
		return "", ErrorNotLogin
	}
	return userName.(string), nil
}
