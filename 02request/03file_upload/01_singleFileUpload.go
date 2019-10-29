package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//单文件上传

func main() {
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		fmt.Println("fileName=", file.Filename)

		//保存到指定目录
		err := c.SaveUploadedFile(file, "D:/test/"+file.Filename)
		if err != nil {
			c.String(http.StatusOK, fmt.Sprintf("'%v' upload fail!", file.Filename))
		}

		c.String(http.StatusOK, fmt.Sprintf("'%v' upload success!", file.Filename))
	})

	router.Run()
}
