package entity

import "Todolist/app/orm"

type CustomItem struct {
	Id   int `json:"id"      params:"type: int, primary_key: true, auto_increment: true"`
	Item int `json:"item"    params:"type: int, foreign_key: item"`
}

func (c *CustomItem) GetItem() *Item {
	return orm.FindById(&Item{}, c.Item).(*Item)
}
