package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//定义绑定结构体
type Login struct {
	User     string `form:"users" json:"users" xml:"users" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

//模型绑定
func main() {
	router := gin.Default()

	//绑定Json({"users": "manu", "password": "123"})
	router.POST("loginJson", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBindJSON(&login); err != nil {
			log.Fatalln(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if login.User != "song" || login.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	//绑定Xml
	//	<?xml version="1.0" encoding="UTF-8"?>
	//	<root>
	//		<users>users</users>
	//		<password>123</password>
	//	</root>
	router.POST("/loginXml", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBindXML(&login); err != nil {
			log.Fatalln(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if login.User != "song" || login.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	//帮定表单(users=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBindQuery(&login); err != nil {
			log.Fatalln(err)
			c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
			return
		}
		if login.User != "song" || login.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.Run(":8080")
}
