package entity

import "Todolist/app/orm"

type Server struct {
	Id       int    `json:"id"       params:"type: int, primary_key: true, auto_increment: true"`
	Name     string `json:"name"     params:"type: varchar, length: 255"`
	Language string `json:"language" params:"type: varchar, length: 255"`
}

func (s *Server) GetPortals() []*Portal {
	return orm.FindByField(&Portal{}, "server", s.Id).([]*Portal)
}
