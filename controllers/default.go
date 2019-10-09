package controllers

import (
	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	"bj4q/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) ShowTest() {
	c.Data["k"] = "路由测试"
	c.TplName = "test.html"
}
func (c *MainController) ShowPost() {
	/*orm:=orm.NewOrm()
	user:=new(models.User)
	user.UserName="张三"
	user.Pwd="FFhu799975844"
	orm.Insert(user)
	c.Data["luyou"]="这是post请求"
	c.TplName="test.html"*/
/*	o := orm.NewOrm()
	user := new(models.User)
	user.Id = 5
	o.Read(user,"")
	fmt.Println(*user)
	c.Data["luyou"]="这是post请求"
	c.TplName="test.html"*/
/*	o:=orm.NewOrm()
	var user models.User
	user.UserName="张三"
	o.Read(&user,"UserName")
	user.Pwd="654321"
	o.Update(&user)
	c.Data["luyou"]="这是post请求"
	c.TplName="test.html"*/
	o:=orm.NewOrm()
	user:=new(models.User)
	user.Id=2
	o.Delete(user)
	c.Data["luyou"]="这是post请求"
	c.TplName="test.html"

	}
