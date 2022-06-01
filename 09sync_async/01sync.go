package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 同步、异步
func main() {
	r := gin.Default()
	r.GET("/sync", func(c *gin.Context) {
		fmt.Println("同步请求")
		time.Sleep(time.Second * 2)
		c.String(http.StatusOK, "同步处理中...")
	})
	r.GET("/async", func(c *gin.Context) {
		ct := c.Copy()
		go func(*gin.Context) {
			time.Sleep(time.Second * 2)
			fmt.Println("异步请求")
		}(ct)
		c.String(http.StatusOK, "异步处理中...")
	})

	r.Run()
}
