package routers

import (
	"github.com/astaxie/beego"
	"go-pest-game/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/user/list", &controllers.UserController{}, "get:List")
}
