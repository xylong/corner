package models

import "strings"

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
	fieldStr := "b." + strings.Join(fields, "b.")
}
