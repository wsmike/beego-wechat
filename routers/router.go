package routers

import (
	"hello/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/hello", &controllers.MainController{})	
    beego.Router("/",&controllers.MainController{})
    beego.Router("/all/:key", &controllers.MainController{}, "get:AllBlock")
    beego.Router("/allPost", &controllers.MainController{}, "post:AllPost")
    beego.Router("/user", &controllers.UserController{}, "get:GetUser")
    beego.Router("/ws", &controllers.WsController{})
}
