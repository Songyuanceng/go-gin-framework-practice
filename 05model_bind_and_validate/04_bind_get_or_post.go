package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Student struct {
	Name     string `form:"name"`
	Address  string `form:"address"`
	Birthday string `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

//绑定get或者post请求参数
func main() {
	route := gin.Default()
	//route.GET("/testing", start)
	route.POST("/testing", start)
	route.Run(":8080")
}

func start(c *gin.Context) {
	var student Student
	if c.ShouldBind(&student) == nil {
		log.Println("不管是post还是get请求,数据绑定成功")
		log.Println(student.Name)
		log.Println(student.Address)
		log.Println(student.Birthday)
	}
	c.JSON(http.StatusOK, "success")
}
