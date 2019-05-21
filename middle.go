package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"log"
)

func main()  {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		go func(c *gin.Context) {
			time.Sleep(5 * time.Second)
			log.Println("done in path " + c.Request.URL.Path)
		}(c)
	})

	r.GET("/long_sync", func(c *gin.Context) {
		// 用 time.Sleep() 模拟一个长任务。
		time.Sleep(5 * time.Second)

		// 因为没有使用 goroutine，不需要拷贝上下文
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	r.Run()
}
