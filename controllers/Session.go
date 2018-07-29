package controllers

import (
	"github.com/astaxie/beego"
	"HM/beego1/demo/loveHome/utils"
	"HM/beego1/demo/loveHome/models"
)

type SessionController struct {
	beego.Controller
}
func (this *SessionController)RetData(resp map[string]interface{}){
	this.Data["json"] = resp
	this.ServeJSON()
}
//登陆账号
func (this *SessionController)GetSessionData(){
	beego.Info("GetSession success")
	resp := make(map[string]interface{})
	defer this.RetData(resp) //数据转成json格式发送给前端


	resp["errno"] = utils.RECODE_OK
	resp["errMsp"] = utils.RecodeText(utils.RECODE_OK)

	user :=models.User{}
	name := this.GetSession("name")
	if name != nil{
		user.Name = name.(string)
		resp["errno"] = utils.RECODE_OK
		resp["errMsp"] = utils.RecodeText(utils.RECODE_OK)
		resp["data"] = user
	}
		resp["errno"] = utils.RECODE_OK
		resp["errMsp"] = utils.RecodeText(utils.RECODE_OK)
		resp["data"] = user

}

//退出登陆
func(this *SessionController)DelSessionData(){
	resp := make(map[string]interface{})
	defer this.RetData(resp)

	resp["errno"] = utils.RECODE_OK
	resp["errMsp"] = utils.RecodeText(utils.RECODE_OK)

	user := models.User{}
	this.DelSession("name")

	resp["errno"] = utils.RECODE_OK
	resp["errMsp"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = user

}
