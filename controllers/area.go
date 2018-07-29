package controllers

import (
	"github.com/astaxie/beego"
	"HM/beego1/demo/loveHome/models"
	"github.com/astaxie/beego/orm"
	"HM/beego1/demo/loveHome/utils"
)

type AreaController struct {
	beego.Controller
}

func (this *AreaController)RetData(resp map[string]interface{}){
	this.Data["json"] = resp
	this.ServeJSON()
}

func (c *AreaController) GetArea() {
	beego.Info("connect success")
	resp := make(map[string]interface{})
	defer c.RetData(resp)  //所有的程序执行完，都会转成json格式

	var areas []models.Area
	//var area = models.Area{}
	o := orm.NewOrm()
	//err := o.Read(&area)//查询表格中的单条数据
	//查询表格中的所有数据
	num,err := o.QueryTable("area").All(&areas)
	if err != nil{
		resp["errNo"] = utils.RECODE_DBERR
		resp["errMsp"] = utils.RecodeText(utils.RECODE_DBERR)

		return
	}
	if num == 0{
		resp["errNo"] = utils.RECODE_NODATA
		resp["errMsp"] = utils.RecodeText(utils.RECODE_NODATA)

		return
	}

	resp["errNo"] = utils.RECODE_OK
	resp["errMsp"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = areas

	beego.Info("query data sucess,resp=",resp,"num =",num)
}
