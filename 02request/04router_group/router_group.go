package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//请求路由分组

func main() {
	router := gin.Default()

	group_v1 := router.Group("/users")
	{
		group_v1.POST("/register", registering)
		group_v1.POST("/login", logining)
		group_v1.POST("/logOut", logOuting)
	}

	group_v2 := router.Group("/manager")
	{
		group_v2.POST("/login", logining)
		group_v2.POST("/logOut", logOuting)
	}

	router.Run(":8080")

}

func registering(c *gin.Context) {
	name := c.PostForm("name")
	age := c.PostForm("age")
	log.Println("name:", name, "age:", age)

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "register success",
	})
}

func logining(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "login success",
	})
}

func logOuting(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "logOut success",
	})
}
