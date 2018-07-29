package main

import (
	_ "HM/beego1/demo/loveHome/routers"
	"github.com/astaxie/beego"
	"strings"
	"net/http"
	"github.com/astaxie/beego/context"
	_"HM/beego1/demo/loveHome/models"
)

func main() {
	ignoreStaticPath()

	beego.Run()
	//beego.Run(":8899") //此处也能设置端口，并且是优先使用
}

func ignoreStaticPath(){
	//透明statci
	beego.InsertFilter("/",beego.BeforeRouter,TransparentStatic)
	beego.InsertFilter("/*",beego.BeforeRouter,TransparentStatic)
}

func TransparentStatic(ctx *context.Context){
	orpath := ctx.Request.URL.Path
	beego.Debug("request url:",orpath)
	//如果请求url还有api字段，说明指令应该 取消静态资源路径重定向

	if strings.Index(orpath,"api") >= 0{
		return
	}

	http.ServeFile(ctx.ResponseWriter,ctx.Request,"static/html"+ctx.Request.URL.Path)  //默认返回主网址
}
