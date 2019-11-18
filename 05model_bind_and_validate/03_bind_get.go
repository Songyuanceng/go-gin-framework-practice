package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

//仅仅绑定get请求参数
func main() {
	route := gin.Default()
	route.Any("/testing", startPage)
	route.Run(":8080")
}

func startPage(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("仅绑定get请求参数成功")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(http.StatusOK, "success")
}
