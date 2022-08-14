package entity

import "Todolist/app/orm"

type CustomEffect struct {
	Id     int `json:"id"          params:"type: int, primary_key: true, auto_increment: true"`
	Effect int `json:"effect"      params:"type: int, foreign_key: effect"`
	Min    int `json:"min"         params:"type: int"`
	Max    int `json:"max"         params:"type: int"`
}

func (c *CustomEffect) GetEffect() *Effect {
	return orm.FindById(&Effect{}, c.Effect).(*Effect)
}
