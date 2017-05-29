package main

import (
	"excel/models"
	_ "excel/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("username")+":"+beego.AppConfig.String("password")+"@/paging?charset=utf8&loc=Asia%2FShanghai", 30)
	orm.RegisterModel(new(models.Person))
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
