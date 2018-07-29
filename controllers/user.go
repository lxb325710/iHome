package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"HM/beego1/demo/loveHome/models"
	"HM/beego1/demo/loveHome/utils"
)

type UserController struct{
	beego.Controller
}

func (this *UserController)RetData(resp map[string]interface{}){
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *UserController)PostRegUser(){
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &resp)
	if err != nil{
		resp["errN0"] = 4001
		resp["errMsp"] = "接收数据错误"

		return
	}
	//beego.Info("mobil=",resp["mobile"])
	//接收数据，插入数据库
	user := models.User{}
	user.Name = resp["mobile"].(string)
	user.Password_hash = resp["password"].(string)
	user.Mobile = resp["mobile"].(string)

	o := orm.NewOrm()
	id, err := o.Insert(&user)
	if err != nil{
		resp["errN0"] = utils.RECODE_USERERR
		resp["errMsp"] = utils.RecodeText(utils.RECODE_USERERR)
	}

	beego.Info("user_Id =",id)
	resp["errN0"] = utils.RECODE_OK
	resp["errMsp"] = utils.RecodeText(utils.RECODE_OK)

	this.SetSession("name",user.Name)

}
