package routers

import (
	"HM/beego1/demo/loveHome/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/api/v1.0/areas", &controllers.AreaController{},"get:GetArea")
    beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexController{},"get:GetHouseIndex")
    beego.Router("/api/v1.0/session", &controllers.SessionController{},"get:GetSessionData;delete:DelSessionData")

    beego.Router("/api/v1.0/sessions", &controllers.UserController{},"post:PostRegUser")
}
