package entity

import (
	"Todolist/app/orm"
	"database/sql"
)

type TaskGroup struct {
	Id          int            `json:"id"          params:"type: int, primary_key: true, auto_increment: true"`
	Icon        string         `json:"icon"        params:"type: varchar, length: 255, nullable: false"`
	Name        string         `json:"name"        params:"type: varchar, length: 255, nullable: false"`
	Description sql.NullString `json:"description" params:"type: text, nullable: true"`
	CustomId    string         `json:"custom_id"   params:"type: varchar, length: 255, nullable: false"`
	Color       string         `json:"color"       params:"type: varchar, length: 255, nullable: false"`
}

func (g *TaskGroup) GetTasks() []*Task {
	values := orm.FindAllByField(Task{}, "group", g.Id)

	tasks := make([]*Task, len(values))
	for i, v := range values {
		tasks[i] = v.(*Task)
	}

	return tasks
}
