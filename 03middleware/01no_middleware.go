package main

import "github.com/gin-gonic/gin"

//无中间件启动
func main() {
	router := gin.New()

	router.Run(":8081")
}
