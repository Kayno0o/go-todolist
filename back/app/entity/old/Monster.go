package entity

import "Todolist/app/orm"

type Monster struct {
	Id       int    `json:"id"        params:"type: int, primary_key: true"`
	MinLevel int    `json:"min_level" params:"type: int"`
	MaxLevel int    `json:"max_level" params:"type: int"`
	Type     int    `json:"type"      params:"type: int, foreign_key: monster_type"`
	Name     string `json:"name"      params:"type: varchar, length: 255, length: 255"`
}

func (m *Monster) GetType() *MonsterType {
	return orm.FindById(&MonsterType{}, m.Type).(*MonsterType)
}
