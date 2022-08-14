package entity

import "Todolist/app/orm"

type UserTaskGroup struct {
	Id         int `json:"id"    params:"type: int, primary_key: true, auto_increment: true"`
	User       int `json:"user"  params:"type: int, nullable: false, foreign_key: user"`
	Group      int `json:"task"  params:"type: int, nullable: false, foreign_key: task_group"`
	Permission int `json:"permission"  params:"type: int, nullable: false"`
}

func (utg *UserTaskGroup) GetUser() *User {
	return orm.FindById(User{}, utg.User).(*User)
}

func (utg *UserTaskGroup) GetTaskGroup() *TaskGroup {
	return orm.FindById(TaskGroup{}, utg.Group).(*TaskGroup)
}
