package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

//多模板文件--将信息渲染到多个模板文件
//简单版
func main() {
	router := gin.Default()
	router.HTMLRender = createMyRender()
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "Html5 Template Engine",
		})
	})
	router.GET("/article", func(c *gin.Context) {
		c.HTML(http.StatusOK, "article", gin.H{
			"title": "Html5 Article Engine",
		})
	})

	router.Run(":8080")
}

func createMyRender() multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()
	renderer.AddFromFiles("index", "./07_html_handle/templates/index/base.html", "./07_html_handle/templates/index/index.html")
	renderer.AddFromFiles("article", "./07_html_handle/templates/article/base.html", "./07_html_handle/templates/article/index.html", "./07_html_handle/templates/article/article.html")
	return renderer
}

//高级版
func main0202() {
	router := gin.Default()
	router.HTMLRender = loadTemplate("./07_html_handle/templates")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "welcome!",
		})
	})

	router.GET("/article", func(c *gin.Context) {
		c.HTML(http.StatusOK, "base.html", gin.H{
			"title": "Html5 Article Engine!",
		})
	})
	router.Run()
}

func loadTemplate(templatePath string) multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()
	indexs, err := filepath.Glob(templatePath + "/index/*.html")
	if err != nil {
		panic(err.Error())
	}

	articles, err := filepath.Glob(templatePath + "/article/*.html")
	if err != nil {
		panic(err.Error())
	}

	/*for _, article := range articles {
		indexsCopy := make([]string, len(indexs))
		copy(indexsCopy, indexs)
		files := append(indexsCopy, article)
		base := filepath.Base(article)
		renderer.AddFromFiles(base, files...)
	}*/

	for _, index := range indexs {
		articleCopy := make([]string, len(articles))
		copy(articleCopy, articles)
		files := append(articleCopy, index)
		renderer.AddFromFiles(filepath.Base(index), files...)
	}
	return renderer
}
