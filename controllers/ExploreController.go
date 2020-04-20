package controllers

import (
	"corner/models"
	"corner/utils"
	"github.com/astaxie/beego"
	"math"
	"strconv"
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

	books, totalCount, err := models.NewBook().HomeData(pageIndex, pageSize, cid)
	if err != nil {
		beego.Error(err)
		e.Abort("404")
	}
	if totalCount > 0 {
		urlSuffix := ""
		if cid > 0 {
			urlSuffix = urlSuffix + "&cid=" + strconv.Itoa(cid)
		}
		html := utils.NewPaginations(4, totalCount, pageSize, pageIndex, urlPrefix, urlSuffix)
		e.Data["PageHtml"] = html
	} else {
		e.Data["PageHtml"] = ""
	}
	e.Data["TotalPages"] = int(math.Ceil(float64(totalCount) / float64(pageSize)))
	e.Data["list"] = books
}
