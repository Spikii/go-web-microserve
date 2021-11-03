package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.4399.com"
	c.Data["Email"] = "740777550@qq.com"
	c.TplName = "index.html"
}
