package routers

import (
	"controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    //beego.Router("/user/login", &controllers.TestController{})

    beego.Router("/user/login", &controllers.LoginController{})
    beego.Router("/user/info", &controllers.GetInfoController{})
    beego.Router("/user/logout", &controllers.LogoutController{})

    {//statistic
		beego.Router("/statistic/activity", &controllers.ActivityController{})
		beego.Router("/statistic/newly", &controllers.NewlyController{})
		beego.Router("/statistic/user", &controllers.UserController{})
	}

}
