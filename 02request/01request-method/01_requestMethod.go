package main

import "github.com/gin-gonic/gin"

//request支持的请求方式

func main() {
	engine := gin.Default()

	engine.GET("/get", getting)
	engine.POST("/post", postting)
	engine.DELETE("/delete", deleting)
	engine.PUT("/put", putting)
	engine.OPTIONS("/option", optioning)
	engine.PATCH("/patch", patching)
	engine.HEAD("/head", heading)

	engine.Run()
}

func getting(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "get success",
	})
}

func postting(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "post success",
	})
}

func putting(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "put success",
	})
}

func deleting(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "delete success",
	})
}

func optioning(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "option success",
	})
}

func patching(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "patch success",
	})
}

func heading(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "head success",
	})
}
