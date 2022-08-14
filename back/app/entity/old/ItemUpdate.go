package entity

import "Todolist/app/orm"

type ItemUpdate struct {
	Id    int    `json:"id"    params:"type: int, primary_key: true, auto_increment: true"`
	Item  int    `json:"item"  params:"type: int, foreign_key: item"`
	Price int    `json:"price" params:"type: int"`
	User  int    `json:"user"  params:"type: int, foreign_key: user"`
	Date  string `json:"date"  params:"type: datetime"`
}

func (i *ItemUpdate) GetItem() *Item {
	return orm.FindById(&Item{}, i.Item).(*Item)
}
