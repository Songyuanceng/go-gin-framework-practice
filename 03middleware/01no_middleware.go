package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//无中间件启动
func main() {
	router := gin.New()

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	router.Run(":8081")
}
