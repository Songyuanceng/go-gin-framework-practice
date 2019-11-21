package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

//html渲染
// 使用LoadHTMLGlob()--加载一个路径下所有templates 或者 LoadHTMLFiles()--加载指定路径下的html文件
func main01() {
	router := gin.Default()
	router.LoadHTMLGlob("./07_html_handle/templates/*")
	//router.LoadHTMLFiles("./07_html_handle/templates/template1.html", "./07_html_handle/templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "This is a website!",
		})

		/*c.HTML(http.StatusOK, "template1.html", gin.H{
			"title": "This is a Main website!",
		})*/
	})
	router.Run(":8080")
}

//渲染不同路径下同名template文件
func main02() {
	router := gin.Default()
	router.Delims("{[{", "}]}") //自定义渲染分隔符
	router.LoadHTMLGlob("07_html_handle/templates/**/*")

	router.GET("/template1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{"title": "this is the first of Template!"})
	})

	router.GET("/template2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods/index.tmpl", gin.H{
			"title":   "this is the second of Tmeplate!",
			"content": "Do you know?"})
	})

	router.Run(":8080")
}

//自定义模板函数
func main() {
	router := gin.Default()
	router.Delims("{", "}")
	router.SetFuncMap(template.FuncMap{"formartAsDate": formartAsDate})
	//router.LoadHTMLGlob("./07_html_handle/**/*")
	router.LoadHTMLFiles("./07_html_handle/templates/raw.tmpl")

	router.GET("/define", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.tmpl", gin.H{
			"now": time.Date(2019, 11, 21, 0, 0, 0, 0, time.UTC)})
	})

	router.Run()
}

func formartAsDate(time time.Time) string {
	year, month, day := time.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}
