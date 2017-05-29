package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

//一页显示的记录数
const PAGESIZE = 3

type Person struct {
	Pid      int    `orm:"pk;auto"`
	NickName string `orm:"size(30);index"`
	UserSex  string `orm:"size(5);null"`
	//用户邮箱
	UserEmail string `orm:"null"`
}

func FindAll() []*Person {
	o := orm.NewOrm()
	allper := make([]*Person, 0)
	o.Raw("select * FROM person").QueryRows(&allper)
	return allper
}

// FindCount 返回符合条件的总记录数
func FindCount() int64 {
	o := orm.NewOrm()
	var count int64
	o.Raw("SELECT COUNT(*) FROM person").QueryRow(&count)
	return count
}

// FindInfo 返回分页的数据
func FindInfo(page int) []*Person {
	o := orm.NewOrm()
	somePel := make([]*Person, 0)
	o.Raw("SELECT * FROM person ORDER BY pid DESC LIMIT ?,?", page, PAGESIZE).QueryRows(&somePel)

	return somePel
}

func FindInofoTwo(offset, limit int) []*Person {
	o := orm.NewOrm()
	somePel := make([]*Person, 0)
	o.Raw("SELECT * FROM person ORDER BY pid DESC LIMIT ?,?", offset, limit).QueryRows(&somePel)

	return somePel
}

func Update(users []Person)(string, error)  {
	o:=orm.NewOrm()
	err:= o.Begin()
	if err!=nil {
		beego.Error("开启事务失败")
		return "cannot open tx",err
	}
	
	for _,val := range users {
		// 这里面可以都数据进行数据格式验证,验证不通过数据回滚,这里我们就不验证了
		
		_,err:= o.Insert(&val)
		if err!=nil {
			o.Rollback()
			
		}
	}

	err=o.Commit()
	if err!=nil {
		beego.Error("添加数据库失败")
		return "update fail",err
	}
	return "update ok",nil
}
