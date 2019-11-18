package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type teacher struct {
	Name string `uri:"name" binding:"required"`
	Age  string `uri:"age" binding:"required"`
}

//绑定URI上的参数/test/:name/:age
func main() {
	router := gin.Default()
	router.GET("/test/:name/:age", testUri)
	router.Run(":8080")
}

func testUri(c *gin.Context) {
	var teacher teacher
	if err := c.ShouldBindUri(&teacher); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	} else {
		log.Println("绑定URI参数成功")
		log.Println(teacher.Name)
		log.Println(teacher.Age)
	}
	c.String(http.StatusOK, "success")
}
