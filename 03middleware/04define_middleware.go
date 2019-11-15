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

		ctx.Next() //传递到下一个func(gin.Context)

		//after request
		latency := time.Since(t)
		log.Print(latency)

		//access the status we are sending
		status := ctx.Writer.Status()
		log.Println(status)
	}

}

func main() {
	r := gin.New()
	r.Use(logger())

	r.GET("/test", func(ctx *gin.Context) {
		example := ctx.MustGet("example").(string)

		//it would print: "12345"
		log.Println(example)
	})

	//Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
