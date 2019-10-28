package main

import "github.com/gin-gonic/gin"

//gin-web-framework框架helloworld
func main() {
	//engine := gin.New() //无中间件启动
	engine := gin.Default() //默认中间件启动 ogger and recovery (crash-free) 中间件

	engine.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "success",
		})
	})

	//engine.Run() //默认8080
	engine.Run(":8888")
}
