package main

import (
	_ "hello/routers"
	"hello/models"
	"github.com/astaxie/beego"
)



func main() {
	models.Init()
	beego.Run()

}
