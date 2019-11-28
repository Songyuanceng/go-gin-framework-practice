package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//支持Let's Encrypt证书
func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

}
