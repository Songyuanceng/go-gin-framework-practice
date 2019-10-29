package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//多文件上传

func main() {
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		multipartFile, err := c.MultipartForm()
		if err != nil {
			log.Fatalln("发生错误：", err)
		}
		files := multipartFile.File["upload[]"] //postman测试的(在key字段写upload[])

		for _, file := range files {
			log.Println("fileName=", file.Filename)
			//存放到指定目录
			c.SaveUploadedFile(file, "D:/test/"+file.Filename)
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "upload success",
		})
	})

	router.Run()

}
