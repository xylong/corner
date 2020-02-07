package models

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
