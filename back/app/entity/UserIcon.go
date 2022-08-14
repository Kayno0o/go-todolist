package entity

import "Todolist/app/orm"

type UserIcon struct {
	Id    int    `json:"id"   params:"type: int, primary_key: true, auto_increment: true"`
	User  int    `json:"user" params:"type: int, nullable: false, foreign_key: user"`
	Icon  string `json:"icon"      params:"type: varchar, length: 255, nullable: false"`
	Color string `json:"color" params:"type: varchar, length: 255, nullable: false"`
}

func (i *UserIcon) GetUser() *User {
	return orm.FindById(User{}, i.User).(*User)
}
