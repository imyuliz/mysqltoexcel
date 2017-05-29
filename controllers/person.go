package controllers

import (
	m "excel/models"
	"fmt"
	"os"
	"path/filepath"

	"github.com/astaxie/beego"

	"strconv"

	"github.com/tealeg/xlsx"
)

type PersonController struct {
	BaseController
	upload chan int
}

func (li *PersonController) AllPeople() {
	//每页显示的记录数
	limit := 3
	//查询符合条件的总记录数
	count := m.FindCount()

	pager := li.SetPaginator(limit, count)
	person := m.FindInofoTwo(pager.Offset(), limit)
	li.Data["Allp"] = person
	li.TplName = "page.html"
}

// Download 把数据库的信息下载到excel中
func (li *PersonController) Download() {

	// li.TplName="index.html"

	li.Ctx.Request.ParseForm()
	// ctx.Request.ParseForm()
	//设置响应头
	li.Ctx.ResponseWriter.Header().Set("Content-Type", "application/vnd.ms-excel")

	filename := filepath.Join(os.Getenv("GOPATH"), "src", "excel", "views", "UserInfoTemplate.xlsx")
	fmt.Println(3)
	beego.Info(filename)

	file, err := xlsx.OpenFile(filename)
	if err != nil {
		//这里的错误处理应该打印日志后跳转
		beego.Error("cann't open file", err)
		li.Redirect("/user/all", 302)
		return
	}
	//设置工作表名
	sheet, ok := file.Sheet["userinfo"]
	if !ok {
		beego.Error("cann't find userinfo sheet")
		//这里的错误处理应该打印日志后跳转
		li.Redirect("/user/all", 302)
		return
	}
	//查询数据库
	per := m.FindAll()
	for _, v := range per {
		row := sheet.AddRow()

		//用户编号
		cell1 := row.AddCell()
		cell1.Value = strconv.Itoa(v.Pid)
		cell1.SetStyle(sheet.Rows[1].Cells[0].GetStyle())

		//用户昵称
		cell2 := row.AddCell()
		cell2.Value = v.NickName
		cell2.SetStyle(sheet.Rows[1].Cells[1].GetStyle())
		//用户性别
		cell3 := row.AddCell()
		cell3.Value = v.UserSex
		cell3.SetStyle(sheet.Rows[1].Cells[2].GetStyle())
		//用户email
		cell4 := row.AddCell()
		cell4.Value = v.UserEmail
		cell4.SetStyle(sheet.Rows[1].Cells[3].GetStyle())

	}

	if len(sheet.Rows) >= 3 {
		sheet.Rows = append(sheet.Rows[0:1], sheet.Rows[2:]...)

	}

	file.Write(li.Ctx.ResponseWriter)

}
