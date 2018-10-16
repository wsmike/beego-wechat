package controllers

import (
	"github.com/astaxie/beego"
    "hello/models"
)




type UserController struct {
	beego.Controller
}


func (this *UserController) GetUser() {

	user, _ := models.GetUserById(1)
    this.Data["json"] = user  
	this.ServeJSON()

}
