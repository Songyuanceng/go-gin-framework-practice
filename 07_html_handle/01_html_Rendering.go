package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//html渲染 使用LoadHTMLGlob() 或者 LoadHTMLFiles()
func main() {
	router := gin.Default()
	//router.LoadHTMLGlob("./07_html_handle/templates/*")
	router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "template1.html", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")
}
