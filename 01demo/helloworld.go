package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//gin-web-framework框架helloworld
func main() {
	//engine := gin.New() //无中间件启动
	engine := gin.Default() //默认中间件启动 logger and recovery (crash-free) 中间件

	engine.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "success",
		})
	})

	//四种启动方式
	//engine.Run() //默认8080
	//engine.Run(":8888") //指定端口
	//http.ListenAndServe(":8080", engine) //go类库提供的方法, engine实现了Handler接口

	s := &http.Server{
		Addr:           ":8080",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe() //自定义服务器配置

}
