package entity

import "Todolist/app/orm"

type Auction struct {
	Id        int    `json:"id"          params:"type: int, primary_key: true, auto_increment: true"`
	Item      int    `json:"item"        params:"type: int, foreign_key: custom_item"`
	User      int    `json:"user"        params:"type: int, foreign_key: user"`
	Price     int    `json:"price"       params:"type: int"`
	IsActive  bool   `json:"is_active"   params:"type: tinyint"`
	CreatedAt string `json:"created_at"  params:"type: datetime"`
	EndsOn    string `json:"ends_on"     params:"type: datetime"`
}

func (a *Auction) GetItem() *CustomItem {
	return orm.FindById(&CustomItem{}, a.Item).(*CustomItem)
}

func (a *Auction) GetUser() *User {
	return orm.FindById(&User{}, a.User).(*User)
}
