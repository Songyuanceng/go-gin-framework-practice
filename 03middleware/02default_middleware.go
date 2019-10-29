package main

import "github.com/gin-gonic/gin"

//默认中间件

func main() {
	//Logger()和Recovery()
	router := gin.Default()

	router.Run(":8081")
}
