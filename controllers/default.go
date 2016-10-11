package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Default() {
	c.Data["webtitle"] = "**官网"
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "templet.html"
}
func (c *MainController) Index() {
	c.TplName = "index.html"
}
