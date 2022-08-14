package entity

import "Todolist/app/orm"

type Almanax struct {
	Id          int    `json:"id"          params:"type: int, primary_key: true"`
	Item        int    `json:"item"        params:"type: int, foreign_key: item"`
	Name        string `json:"name"        params:"type: varchar, length: 255"`
	Description string `json:"description" params:"type: text"`
	Npc         string `json:"npc"         params:"type: varchar, length: 255"`
	Date        string `json:"date"        params:"type: date"`
}

func (a *Almanax) GetItem() *Item {
	return orm.FindById(&Item{}, a.Item).(*Item)
}
