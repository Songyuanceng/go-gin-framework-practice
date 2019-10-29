package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取get和post混合请求方式参数

func main() {
	router := gin.Default()

	router.POST("/user", func(c *gin.Context) {
		id := c.Query("id")
		age := c.DefaultQuery("age", "0") //从url中获取

		name := c.PostForm("name")
		message := c.PostForm("message") //从form表单中获取

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, age, name, message)
		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"name":    name,
			"age":     age,
			"message": message,
		})
	})

	router.Run()
}
