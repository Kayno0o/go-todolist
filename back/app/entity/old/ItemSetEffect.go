package entity

import "Todolist/app/orm"

type ItemSetEffect struct {
	Id           int `json:"id"            params:"type: int, primary_key: true, auto_increment: true"`
	ItemNumber   int `json:"item_number"   params:"type: int, foreign_key: item"`
	ItemSet      int `json:"item_set"      params:"type: int, foreign_key: item_set"`
	CustomEffect int `json:"custom_effect" params:"type: int, foreign_key: custom_effect"`
}

func (i *ItemSetEffect) GetItemSet() *ItemSet {
	return orm.FindById(&ItemSet{}, i.ItemSet).(*ItemSet)
}

func (i *ItemSetEffect) GetCustomEffect() *CustomEffect {
	return orm.FindById(&CustomEffect{}, i.CustomEffect).(*CustomEffect)
}
