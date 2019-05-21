package middleware

import (
	"github.com/gin-gonic/gin"
	"fmt"
		"github.com/ninorain22/gin_starter/common"
	"github.com/ninorain22/gin_starter/enumeration"
)

func TokenValidRequired(c *gin.Context) {
	userId := c.Query("userId")
	token := c.Query("token")
	fmt.Println(userId, token)

	// todo: 用户token校验逻辑
	if userId != token {
		common.SendErrJSON("invalid token", enumeration.TOKEN_INVALID, c)
		return
	}
	c.Next()
}