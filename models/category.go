package models

import (
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id     int
	Pid    int
	Title  string `orm:"size(30);unique"`
	Intro  string
	Icon   string
	Cnt    int // 分类统计
	Sort   int
	Status bool // 显示隐藏
}

func (c *Category) TableName() string {
	return TNCategory()
}

// GetCates 获取分类
func (c *Category) GetCategories(pid, status int) (categories []Category, err error) {
	qs := orm.NewOrm().QueryTable(c.TableName())
	if pid > -1 {
		qs = qs.Filter("pid", pid)
	}
	if status == 0 || status == 1 {
		qs = qs.Filter("status", status)
	}
	_, err = qs.OrderBy("-status", "sort", "title").All(&categories)
	return
}

// Find 获取单个分类
func (c *Category) Find(cid int) (category Category) {
	category.Id = cid
	orm.NewOrm().Read(&category)
	return
}
