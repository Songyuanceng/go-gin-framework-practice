package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取get请求参数

func main() {
	router := gin.Default()

	// 匹配的url格式:  /user?firstname=Jane&lastname=Doe
	router.GET("/user", func(c *gin.Context) {
		firstName := c.DefaultQuery("firstname", "song") //带有默认值获取参数,获取不到firstname值时默认song
		lastName := c.Query("lastname")                  //c.Request.URL.Query().Get("lastname")的简写

		c.String(http.StatusOK, "hello %s %s", firstName, lastName)
	})

	router.Run()
}
