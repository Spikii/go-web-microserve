package controllers

import (
	"beego/dao"

	beego "github.com/beego/beego/v2/server/web"
)

type ClassController struct {
	beego.Controller
}

func (ctx *ClassController) Get() {
	class, _ := dao.GetClasses()
	ctx.Data["json"] = class
	ctx.ServeJSON()
}

func (ctx *ClassController) Test() {
	class, _ := dao.GetClasses()
	ctx.Data["json"] = class
	ctx.ServeJSON()
}
