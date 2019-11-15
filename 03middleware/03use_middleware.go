package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// 创建一个不包含中间件的路由器
	r := gin.New()

	// 全局中间件
	// 使用 Logger 中间件
	r.Use(gin.Logger())

	// 使用 Recovery 中间件
	r.Use(gin.Recovery())

	// 路由添加中间件，可以添加任意多个
	r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// 路由组中添加中间件
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(AuthRequired())
	{
		authorized.POST("/register", registering)
		authorized.POST("/login", logining)
		authorized.POST("/logout", logOuting)

		// nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

func MyBenchLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
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

func analyticsEndpoint(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "analytics success",
	})
}

func benchEndpoint(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "benchEndpoint success",
	})
}
