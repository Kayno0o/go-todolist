package entity

type ItemUpdateVote struct {
	Id         int  `json:"id"          params:"type: int, primary_key: true, auto_increment: true"`
	User       int  `json:"user"        params:"type: int, foreign_key: user"`
	ItemUpdate int  `json:"item_update" params:"type: int, foreign_key: item_update"`
	Upvote     bool `json:"upvote"      params:"type: tinyint"`
}
