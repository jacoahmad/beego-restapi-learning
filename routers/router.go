package routers

import (
	"github.com/jacoahmad/learning-bee-go/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
