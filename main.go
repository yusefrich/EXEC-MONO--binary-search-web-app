package main

import (
	_ "helloWeb/routers"

	"github.com/astaxie/beego"
)

//template functions

func main() {
	beego.SetStaticPath("/views/css", "css")
	beego.Run()
}
