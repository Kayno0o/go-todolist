package entity

import "Todolist/app/orm"

type ArchMonster struct {
	Id    int    `json:"id"    params:"type: int, primary_key: true"`
	Price int    `json:"price" params:"type: int"`
	Type  int    `json:"type"  params:"type: int, foreign_key: monster_type"`
	Name  string `json:"name"  params:"type: varchar, length: 255, length: 255"`
}

func (a *ArchMonster) GetType() *MonsterType {
	return orm.FindById(&MonsterType{}, a.Type).(*MonsterType)
}
