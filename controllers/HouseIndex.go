package controllers

import (
	"github.com/astaxie/beego"
	"HM/beego1/demo/loveHome/utils"
)

type HouseIndexController struct {
	beego.Controller
}

func (this *HouseIndexController)RetData(resp map[string]interface{}){
	this.Data["json"] = resp
	this.ServeJSON()
}

func (c *HouseIndexController)GetHouseIndex(){
	resp := make(map[string]interface{})
	defer c.RetData(resp)

	resp["errNo"] = utils.RECODE_DBERR
	resp["errMsp"] = utils.RecodeText(utils.RECODE_DATAERR)

}