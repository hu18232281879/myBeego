package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)
type User struct {
	Id int
	UserName string
	Pwd string
}
func init() {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/t1")
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default",false,true)
}
