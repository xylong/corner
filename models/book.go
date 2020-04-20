package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)

type Book struct {
	Id          int
	Name        string `orm:"size(500)"json:"name"`
	Identify    string `orm:"size((100);unique;"json:"identify"`
	OrderIndex  int    `orm:"default(0)"json:"order_index"`
	Description string `orm:"size(1000)"json:"description"`
}

func (b *Book) TableName() string {
	return TNBook()
}

func NewBook() *Book {
	return &Book{}
}

func (b *Book) HomeData(pageIndex, pageSize, cid int, fields ...string) (books []Book, totalCount int, err error) {
	if len(fields) == 0 {
		fields = append(fields, "book_id", "book_name", "identify", "cover", "order_index")
	}
	fieldStr := "b." + strings.Join(fields, ",b.")
	sqlFormat := "select %v from %s b left join %s c on b.book_id=c.book_id where c.category.id=" + strconv.Itoa(cid)
	sql := fmt.Sprintf(sqlFormat, fieldStr, b.TableName(), TNBookCategory())
	sqlCount := fmt.Sprintf(sqlFormat, "count(*) cnt")

	o := orm.NewOrm()
	var params []orm.Params
	if _, err := o.Raw(sqlCount).Values(&params); err != nil {
		if len(params) > 0 {
			totalCount, _ = strconv.Atoi(params[0]["cnt"].(string))
		}
	}
	_, err = o.Raw(sql).QueryRows(&books)
	return
}
