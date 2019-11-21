package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取post请求参数

func main() {
	router := gin.Default()

	router.POST("/users", func(c *gin.Context) {
		name := c.PostForm("name")
		age := c.DefaultPostForm("age", "18") //设置默认值

		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	router.Run()
}
