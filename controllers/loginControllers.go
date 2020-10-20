package controllers

import (
	"DataCertPlatform/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginControllers struct {
	beego.Controller
}
/*
直接跳转展示用户登录页面
 */func (l *LoginControllers) Get()  {
	l.TplName = "login.html"
}


func (l *LoginControllers) Post() {
	//1、解析客户端用户提交的登录数据
	var user models.User
	err :=l.ParseForm(&user)
	if err != nil {
		//fmt.Println(err.Error())   可返回错误，用于反查
		l.Ctx.WriteString("抱歉，用户登录信息解析失败，请重新尝试")
		return
	}

	//2、根据解析到的数据，执行数据库查询操作
	u, err := user.QueryUser()

	//3、判断数据库查询结果
	if err != nil {
		fmt.Println(err.Error())   //可返回错误，用于反查
		l.Ctx.WriteString("抱歉，用户登录失败")
		return
	}
	//4、根据查询结果返回客户端相应的信息或者页面跳转
	//l.TplName = "login.html"
	l.Data["Phone"] = u.Phone//动态数据设置
	l.TplName = "home.html"
}


