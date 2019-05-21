package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ninorain22/gintest/model"
	"github.com/ninorain22/gintest/common"
	"github.com/ninorain22/gintest/manager"
	"strconv"
	"github.com/ninorain22/gintest/enumeration"
	)

func Add(c *gin.Context) {
	u := model.User{}
	if err := c.ShouldBind(&u); err != nil {
		common.SendErrJSON("invalid params", c)
		return
	}
	if u.Name == "" {
		common.SendErrJSON("name required", c)
		return
	}
	if _, err := manager.Engine.Insert(&u); err != nil {
		common.SendErrJSON("insert user failed", enumeration.INSERT_FAILED, c)
		return
	}
	common.SendJSON(c, u)
}

func Delete(c *gin.Context) {
	u := model.User{}
	if err := c.ShouldBind(&u); err != nil {
		common.SendErrJSON("invalid params", c)
		return
	}
	if u.Id <= 0 {
		common.SendErrJSON("id required", c)
		return
	}
	if _, err := manager.Engine.Id(u.Id).Delete(&u); err != nil {
		common.SendErrJSON("delete user failed", enumeration.DELETE_FAILED, c)
		return
	}
	common.SendJSON(c, nil)
}

func Update(c *gin.Context) {
	u := model.User{}
	if err := c.ShouldBind(&u); err != nil {
		common.SendErrJSON("invalid params", c)
		return
	}
	if _, err := manager.Engine.Id(u.Id).Update(&u); err != nil {
		common.SendErrJSON("update user failed", enumeration.UPDATE_FAILED, c)
		return
	}
	common.SendJSON(c, u)
}

func Get(c *gin.Context) {
	u := model.User{}
	userId, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	if userId <= 0 {
		common.SendErrJSON("id required", c)
		return
	}
	if has, err := manager.Engine.Id(userId).Get(&u); !has || err != nil {
		common.SendErrJSON("find user failed", enumeration.FIND_FAILED, c)
		return
	}
	common.SendJSON(c, u)
}

func List(c *gin.Context) {
	users := make([]model.User, 0)
	if err := manager.Engine.Find(&users); err != nil {
		common.SendErrJSON("find user failed", enumeration.FIND_FAILED, c)
		return
	}
	common.SendJSON(c, users)
}
