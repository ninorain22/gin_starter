package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ninorain22/gin_starter/controller/test"
	"github.com/ninorain22/gin_starter/controller/user"
	"github.com/ninorain22/gin_starter/middleware"
)


func Route(r *gin.Engine)  {
	v1 := r.Group("/v1", middleware.TokenValidRequired)
	{
		// user
		v1.POST("/user/add", user.Add)
		v1.POST("/user/delete", user.Delete)
		v1.POST("/user/update", user.Update)
		v1.GET("/user/get", user.Get)
		v1.GET("/user/list", user.List)
	}

	t := r.Group("/test")
	{
		t.GET("/ping", test.Ping)
	}
}
