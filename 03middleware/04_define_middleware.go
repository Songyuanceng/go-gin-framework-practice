package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		//Set example variable
		ctx.Set("example", "12345")

		//before request

		ctx.Next() //放行，传递到下一个func(gin.Context)
		//ctx.Abort()//退出当前中间件处理

		//after request
		latency := time.Since(t)
		log.Println("time", latency)

		//access the status we are sending
		status := ctx.Writer.Status()
		log.Println("中间件执行完毕", status)
	}

}

func main() {
	r := gin.New()
	r.Use(logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		//it would print: "12345"
		log.Println("example", example)
	})

	//Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
