package entity

type PortalVote struct {
	Id     int  `json:"id"     params:"type: int, primary_key: true, auto_increment: true"`
	Portal int  `json:"portal" params:"type: int, foreign_key: portal"`
	User   int  `json:"user"   params:"type: int, foreign_key: user"`
	Date   int  `json:"date"   params:"type: int, default: 0"`
	Upvote bool `json:"upvote" params:"type: tinyint, default: false"`
}
