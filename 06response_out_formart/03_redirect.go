package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//重定向
func main() {
	router := gin.Default()

	//发布HTTP重定向很容易，支持内部和外部链接
	router.GET("/go", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	//Gin路由重定向，使用如下的HandleContext
	router.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		router.HandleContext(c)
	})
	router.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})
	router.Run(":8080")
}
