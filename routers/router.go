package routers

import (
	"github.com/chenquanquan/golepv/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/client", &controllers.ClientController{})
    beego.Router("/monitor", &controllers.MonitorController{})
}
