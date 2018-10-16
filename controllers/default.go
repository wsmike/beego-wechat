package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)


type MainController struct {
	beego.Controller
}


func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	
}


func (c *MainController) AllBlock() {
	id := c.GetString(":key")
	c.Data["xsrf"]=c.XSRFToken() 
	c.Data["id"]= id
    c.TplName = "admin/admin.tpl"
}


func (this *MainController) AllPost() {
	this.EnableXSRF = false
	// list := &User{1,"yj",20,"m","s"}
    // beego  json 格式
    //    this.Data["json"] = list  
	// this.ServeJSON()
	list := this.Ctx.Request

    fmt.Printf("%d",list)
}
