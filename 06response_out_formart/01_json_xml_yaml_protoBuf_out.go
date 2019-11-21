package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

type name struct {
}

//XML、JSON、YAML和ProtoBuf 渲染（输出格式）
func main() {
	router := gin.Default()

	//输出Json格式
	router.GET("/testJson", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	//输出结构体json
	router.GET("/testJson2", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"users"`
			Message string
			Number  int
		}

		msg.Name = "张三"
		msg.Message = "你去哪了"
		msg.Number = 18

		c.JSON(http.StatusOK, msg)
	})

	//输出Xml
	router.GET("/testXml", func(c *gin.Context) {
		//c.XML(http.StatusOK, gin.H{"message":"success","status":http.StatusOK})
		var info struct {
			Name    string
			Message string
			Number  int
		}

		info.Name = "张三"
		info.Message = "你去哪了"
		info.Number = 18

		c.XML(http.StatusOK, gin.H{"info": info})
	})

	//输出Yaml
	router.GET("/testYml", func(c *gin.Context) {
		var info struct {
			Name    string
			Message string
			Number  int
		}

		info.Name = "张三"
		info.Message = "你去哪了"
		info.Number = 18
		c.YAML(http.StatusOK, info)
		//c.YAML(http.StatusOK, gin.H{"message":"success","status":http.StatusOK})

	})

	//输出ProtoBuf(序列化和反序列化)
	router.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// The specific definition of protobuf is written in the testdata/protoexample file.
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// Note that data becomes binary data in the response
		// Will output protoexample.Test protobuf serialized data
		c.ProtoBuf(http.StatusOK, data)
	})

	//输出secureJson--为防止json被劫持，如果返回的数据是数组，则会默认在返回值前加上"while(1)"
	router.GET("/testSecureJson", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// 将会输出:   while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
		//c.SecureJSON(http.StatusOK, gin.H{"message":"success","status":http.StatusOK})
	})

	//输出asciJson--使用AsciiJSON将使特殊字符编码
	router.GET("/asciJson", func(c *gin.Context) {
		//输出：{"message":"\u6210\u529f","status":200}
		c.AsciiJSON(http.StatusOK, gin.H{"message": "成功", "status": http.StatusOK})
	})

	//输出pureJson
	//通常情况下，JSON会将特殊的HTML字符替换为对应的unicode字符，比如<替换为\u003c，如果想原样输出html，则使用PureJSON
	router.GET("/testPureJson", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{"html": "<b>Hello, world!</b>"}) //输出{"html":"\u003cb\u003eHello, world!\u003c/b\u003e"}
		c.PureJSON(http.StatusOK, gin.H{"html": "<b>Hello, world!</b>"}) //{"html":"<b>Hello, world!</b>"}
	})

	//输出jsonp--使用JSONP可以跨域传输，如果参数中存在回调参数，那么返回的参数将是回调函数的形式
	router.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		// 访问 http://localhost:8081/JSONP?callback=call
		// 将会输出:   call({foo:"bar"})
		c.JSONP(http.StatusOK, data)
	})
	router.Run(":8080")
}
