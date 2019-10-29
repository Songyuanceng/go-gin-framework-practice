package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//当map作为get参数或者post参数时

func main() {
	router := gin.Default()

	router.POST("/user", func(c *gin.Context) {
		ids := c.QueryMap("ids")

		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
	})

	router.Run()
}
