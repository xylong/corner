package controllers

import (
	"corner/models"
	"github.com/astaxie/beego"
)

type HomeController struct {
	BaseController
}

func (h *HomeController) Index() {
	if cates, err := new(models.Category).GetCates(-1, 1); err == nil {
		h.Data["Cates"] = cates
	} else {
		beego.Error(err.Error())
	}
	h.TplName = "home/list.html"
}
