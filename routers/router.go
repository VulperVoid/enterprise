package routers

import (
	"enterprise/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&controllers.MainController{})
}
