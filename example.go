package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"os"
	"io"
	"github.com/gin-gonic/gin/binding"
)

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main()  {
	gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Println(id, page, name, message)
	})

	r.GET("/someJSON", func(c *gin.Context) {
		// gin.H是map[string]interface{}的语法糖
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"status": http.StatusOK,
		})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name string `json:"user"`
			Message string
			Number int
		}
		msg.Name = "Nino"
		msg.Message = "hey"
		msg.Number = 123

		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})

	r.POST("/testPost", func(c *gin.Context) {
		type input struct {
			Id int `json:"id" binding:"required"`
			Name string `json:"name" binding:"required"`
		}
		i := input{}
		if err := c.ShouldBind(&i); err == nil {
			c.JSON(http.StatusOK, i)
		}
		j := input{}
		if err := c.ShouldBindBodyWith(&j, binding.JSON); err == nil {
			c.JSON(http.StatusOK, j)
		}
	})

	auth := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo": "bar",
		"austin": "1234",
	}))

	auth.GET("/secrets", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO"})
		}
	})

	r.Run()
}