package controllers


import (
	"github.com/astaxie/beego"
	
	u "excel/util"
)

type BaseController struct{
	beego.Controller
}
// SetPaginator 设置分页器
func (this *BaseController) SetPaginator(per int, nums int64) *u.Paginator {
	p := u.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}