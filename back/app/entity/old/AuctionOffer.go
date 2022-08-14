package entity

import "Todolist/app/orm"

type AuctionOffer struct {
	Id      int    `json:"id"      params:"type: int, primary_key: true, auto_increment: true"`
	Auction int    `json:"auction" params:"type: int, foreign_key: auction"`
	User    int    `json:"user"    params:"type: int, foreign_key: user"`
	Price   int    `json:"price"   params:"type: int"`
	Created string `json:"created" params:"type: datetime"`
}

func (a *AuctionOffer) GetAuction() *Auction {
	return orm.FindById(&Auction{}, a.Auction).(*Auction)
}

func (a *AuctionOffer) GetUser() *User {
	return orm.FindById(&User{}, a.User).(*User)
}
