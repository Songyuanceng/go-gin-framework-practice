package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

//中间件使用goroutines

func main() {
	router := gin.Default()

	//异步
	router.GET("/testAsync", func(c *gin.Context) {
		copyC := c.Copy()
		go func() {
			time.Sleep(time.Second * 5)
			//使用副本
			log.Println("Done! in path:" + copyC.Request.URL.Path)
		}()
	})

	//同步
	router.GET("/testSync", func(c *gin.Context) {
		time.Sleep(time.Second * 5)
		log.Println("Done! in path:" + c.Request.URL.Path)
	})

	router.Run()
}
