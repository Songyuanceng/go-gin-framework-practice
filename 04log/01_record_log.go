package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	//禁用控制台颜色
	gin.DisableConsoleColor()

	// 创建记录日志的文件
	file, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(file)

	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}
