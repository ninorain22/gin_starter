package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ninorain22/gin_starter/router"
)

func main()  {
	app := gin.Default()

	// 配置路由
	router.Route(app)

	app.Run()
}