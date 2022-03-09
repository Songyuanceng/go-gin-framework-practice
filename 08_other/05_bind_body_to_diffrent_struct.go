package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//绑定请求体到不同结构体
type FormA struct {
	Foo string `form:"foo" xml:"foo" binding:"required"`
}

type FormB struct {
	Bar string `form:"bar" xml:"bar" binding:"required"`
}

func shouldBind(c *gin.Context) {
	foo := FormA{}
	//bar := FormB{}
	if err := c.ShouldBind(&foo); err == nil {
		c.JSON(http.StatusOK, gin.H{})
	}
}
func main() {

}
