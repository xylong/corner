package controllers

import (
	"corner/models"
	"github.com/astaxie/beego"
)

type HomeController struct {
	BaseController
}

func (h *HomeController) Index() {
	if categories, err := new(models.Category).GetCategories(-1, 1); err == nil {
		h.Data["Cates"] = categories
	} else {
		beego.Error(err.Error())
	}
	h.TplName = "home/list.html"
}
