package models

type BookCategory struct {
	Id         int
	BookId     int
	CategoryId int
}

func (b *BookCategory) TableName() string {
	return TNBookCategory()
}
