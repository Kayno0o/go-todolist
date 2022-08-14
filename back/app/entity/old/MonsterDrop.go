package entity

import "Todolist/app/orm"

type MonsterDrop struct {
	Id          int  `json:"id"           params:"type: int, primary_key: true"`
	Monster     int  `json:"monster"      params:"type: int, foreign_key: monster"`
	Item        int  `json:"item"         params:"type: int, foreign_key: item"`
	Chance      int  `json:"chance"       params:"type: int"`
	HasCriteria bool `json:"has_criteria" params:"type: tinyint"`
}

func (m *MonsterDrop) GetItem() *Item {
	return orm.FindById(&Item{}, m.Item).(*Item)
}

func (m *MonsterDrop) GetMonster() *Monster {
	return orm.FindById(&Monster{}, m.Monster).(*Monster)
}
