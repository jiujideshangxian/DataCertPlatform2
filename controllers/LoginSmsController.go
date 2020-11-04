package controllers

import "github.com/astaxie/beego"

type LoginSmsController struct{
	beego.Controller
}

func (l *LoginController) Get(){
	l.TplName = "login_sms.html"
}
