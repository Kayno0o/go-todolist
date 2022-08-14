package entity

import "Todolist/app/orm"

type ItemEffect struct {
	Id           int `json:"id"            params:"type: int, primary_key: true, auto_increment: true"`
	Item         int `json:"item"          params:"type: int, foreign_key: item"`
	CustomEffect int `json:"custom_effect" params:"type: int, foreign_key: custom_effect"`
}

func (i *ItemEffect) GetItem() *Item {
	return orm.FindById(&Item{}, i.Item).(*Item)
}

func (i *ItemEffect) GetCustomEffect() *CustomEffect {
	return orm.FindById(&CustomEffect{}, i.CustomEffect).(*CustomEffect)
}
