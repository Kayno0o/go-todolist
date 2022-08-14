package entity

import "Todolist/app/orm"

type Ingredient struct {
	Id       int `json:"id"       params:"type: int, primary_key: true, auto_increment: true"`
	Craft    int `json:"craft"    params:"type: int, foreign_key: craft"`
	Item     int `json:"item"     params:"type: int, foreign_key: item"`
	Quantity int `json:"quantity" params:"type: int"`
}

func (i *Ingredient) GetCraft() *Craft {
	return orm.FindById(&Craft{}, i.Craft).(*Craft)
}

func (i *Ingredient) GetItem() *Item {
	return orm.FindById(&Item{}, i.Item).(*Item)
}
