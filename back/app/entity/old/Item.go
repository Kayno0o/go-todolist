package entity

import (
	"Todolist/app/orm"
	"database/sql"
)

type Item struct {
	Id          int           `json:"id"          params:"type: int, primary_key: true"`
	ItemType    int           `json:"item_type"   params:"type: int, foreign_key: item_type"`
	Price       int           `json:"price"       params:"type: int, default: 0"`
	Level       int           `json:"level"       params:"type: int"`
	PetXp       int           `json:"pet_xp"      params:"type: int, default: 0"`
	ItemSet     sql.NullInt16 `json:"item_set"    params:"type: int, nullable: true, foreign_key: item_set"`
	Name        string        `json:"name"        params:"type: varchar, length: 255"`
	Description string        `json:"description" params:"type: text"`
}

func (i *Item) GetType() *ItemType {
	return orm.FindById(&ItemType{}, i.ItemType).(*ItemType)
}

func (i *Item) GetSet() *ItemSet {
	if i.ItemSet.Valid {
		return orm.FindById(&ItemSet{}, int(i.ItemSet.Int16)).(*ItemSet)
	}

	return nil
}

func CastItemArray(items []interface{}) []*Item {
	result := make([]*Item, len(items))

	for index, item := range items {
		result[index] = item.(*Item)
	}

	return result
}
