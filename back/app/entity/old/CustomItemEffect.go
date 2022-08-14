package entity

import "Todolist/app/orm"

type CustomItemEffect struct {
	Id           int `json:"id"            params:"type: int, primary_key: true, auto_increment: true"`
	CustomItem   int `json:"custom_item"   params:"type: int, foreign_key: custom_item"`
	CustomEffect int `json:"custom_effect" params:"type: int, foreign_key: custom_effect"`
}

func (c *CustomItemEffect) GetCustomItem() *CustomItem {
	return orm.FindById(&CustomItem{}, c.CustomItem).(*CustomItem)
}

func (c *CustomItemEffect) GetCustomEffect() *CustomEffect {
	return orm.FindById(&CustomEffect{}, c.CustomEffect).(*CustomEffect)
}
