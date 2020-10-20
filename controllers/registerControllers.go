package controllers

import (
	"DataCertPlatform/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Get() {
	r.TplName = "login.html"
}
/*
该方法用于处理用户注册的逻辑
 */
func (r *RegisterController) Post() {
	//1、解析用户端提交的请求数据
	var user models.User
	err :=r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("抱歉，数据解析失败，请重新尝试。")
		return
	}
	//2、将解析到的数据保存到数据库中
	row,err:=models.User.AddUser(user);
	//3、将处理结果返回客户端浏览器
	if err != nil {
		r.Ctx.WriteString("抱歉，注册失败，请重新尝试。")
		return
	}
	fmt.Println(row)


	r.TplName = "login.html"

}
