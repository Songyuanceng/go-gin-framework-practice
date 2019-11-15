package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//获取请求路径中的参数

func main() {
	router := gin.Default()

	// 此规则能够匹配/user/song这种格式，但不能匹配/user/ 或 /user这种格式
	// 示例：localhost/user/song，页面展示: hello song
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello %s", name)
	})

	// 但是，这个规则既能匹配/user/song/格式也能匹配/user/song/upper这种格式
	// 如果没有其他路由器匹配/user/song，它将重定向到/user/song/
	// localhost/user/song/upper，输出：song is upper
	router.GET("/user/:name/*do", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("do")
		do := strings.TrimPrefix(action, "/")
		c.String(http.StatusOK, name+" is "+do)
	})

	router.Run()
}
