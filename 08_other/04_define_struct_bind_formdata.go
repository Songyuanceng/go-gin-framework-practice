package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//使用自定义结构绑定表单数据
type structA struct {
	FieldA string `form:"field_a"`
}

type structB struct {
	NestedStruct structA
	FieldB       string `form:"field_b"`
}

type structC struct {
	NestedStructPointer *structA
	FieldC              string `form:"field_c"`
}

type structD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func getDataB(c *gin.Context) {
	var b structB
	if err := c.ShouldBindQuery(&b); err != nil {
		log.Println("err:", err)
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"a": b.NestedStruct,
			"b": b.FieldB,
		})
}

func getDataC(ctx *gin.Context) {
	var c structC
	if err := ctx.ShouldBindQuery(&c); err != nil {
		log.Println("err:", err)
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"a": c.NestedStructPointer,
			"b": c.FieldC,
		})
}

func getDataD(c *gin.Context) {
	var d structD
	if err := c.ShouldBindQuery(&d); err != nil {
		log.Println("err:", err)
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"a": d.NestedAnonyStruct,
			"d": d.FieldD,
		})
}

func main() {
	router := gin.Default()
	router.GET("/getb", getDataB)
	router.GET("/getc", getDataC)
	router.GET("/getd", getDataD)
	router.Run(":8080")
}
