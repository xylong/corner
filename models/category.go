package models

import "github.com/astaxie/beego/orm"

type Category struct {
	Id     int
	Pid    int
	Title  string `orm:"size(30);unique"`
	Intro  string
	Icon   string
	Cnt    int
	Sort   int
	Status bool // 显示隐藏
}

func (c *Category) TableName() string {
	return TNCategory()
}

// GetCates 获取分类
func (c *Category) GetCates(pid, status int) (cates []Category, err error) {
	qs := orm.NewOrm().QueryTable(c.TableName())
	if pid > -1 {
		qs = qs.Filter("pid", pid)
	}
	if status == 0 || status == 1 {
		qs = qs.Filter("status", status)
	}
	_, err = qs.OrderBy("-status", "sort", "title").All(&cates)
	return
}
