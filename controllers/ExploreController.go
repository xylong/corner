package controllers

import (
	"corner/models"
	"github.com/astaxie/beego"
)

type ExploreController struct {
	BaseController
}

func (e *ExploreController) Index() {
	var (
		cid       int
		category  models.Category
		urlPrefix = beego.URLFor("ExploreController.Index")
	)
	if cid, _ = e.GetInt("cid"); cid > 0 {
		cateModel := new(models.Category)
		category = cateModel.Find(cid)
	}
	e.Data["Cate"] = category
	e.TplName = "explore/index.html"

	pageIndex, _ := e.GetInt("page", 1)
	pageSize := 24
}
