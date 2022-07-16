package controller

import (
	"github.com/gin-gonic/gin"
)

func getCurrentUserName(c *gin.Context) (userName string, err error) {
	_userName, ok := c.Get(ContextUserNameKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userName, ok = _userName.(string)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
