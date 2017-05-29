package routers

import (
	"excel/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/all", &controllers.PersonController{}, "GET:AllPeople")
	beego.Router("/user/download", &controllers.PersonController{}, "GET:Download")
	beego.Router("/user/update", &controllers.PersonController{}, "GET:Update")
	
}
