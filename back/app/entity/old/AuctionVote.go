package entity

import "Todolist/app/orm"

type AuctionVote struct {
	Id      int  `json:"id"      params:"type: int, primary_key: true, auto_increment: true"`
	User    int  `json:"user"    params:"type: int, foreign_key: user"`
	Auction int  `json:"auction" params:"type: int, foreign_key: auction"`
	Upvote  bool `json:"upvote"  params:"type: tinyint"`
}

func (a *AuctionVote) GetUser() *User {
	return orm.FindById(&User{}, a.User).(*User)
}

func (a *AuctionVote) GetAuction() *Auction {
	return orm.FindById(&Auction{}, a.Auction).(*Auction)
}
