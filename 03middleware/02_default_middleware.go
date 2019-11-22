package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//默认中间件

func main() {
	//Logger()和Recovery()
	router := gin.Default()

	router.POST("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	router.Run(":8081")
}
