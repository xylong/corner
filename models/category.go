package models

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
