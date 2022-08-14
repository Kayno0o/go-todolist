package entity

import "Todolist/app/orm"

type ItemType struct {
	Id   int    `json:"id"   params:"type: int, primary_key: true"`
	Name string `json:"name" params:"type: varchar, length: 255"`
}

func (i *ItemType) GetItems() []*Item {
	return CastItemArray(orm.FindAllByField(&Item{}, "item_type", i.Id))
}
