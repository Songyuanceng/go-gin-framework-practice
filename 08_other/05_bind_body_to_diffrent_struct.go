package main

//绑定请求体到不同结构体
type foo struct {
	Foo string `form:"foo" xml:"foo" binding:"required"`
}

func main() {

}
