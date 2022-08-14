package entity

import "Todolist/app/orm"

type Portal struct {
	Id         int    `json:"id"          params:"type: int, primary_key: true, auto_increment: true"`
	UsagesLeft int    `json:"usages_left" params:"type: int, default: 0"`
	X          int    `json:"x"           params:"type: int"`
	Y          int    `json:"y"           params:"type: int"`
	UpdatedAt  int    `json:"updated_at"  params:"type: int, default: 0"`
	Server     int    `json:"server"      params:"type: int, foreign_key: server"`
	Name       string `json:"name"        params:"type: varchar, length: 255"`
}

func (p *Portal) GetServer() *Server {
	return orm.FindById(&Server{}, p.Server).(*Server)
}
