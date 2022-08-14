package entity

import (
	"Todolist/app/orm"
	"database/sql"
)

type Task struct {
	Id        int           `json:"id"         params:"type: int, primary_key: true, auto_increment: true"`
	Amount    sql.NullInt16 `json:"amount"     params:"type: int, nullable: true"`
	Completed int           `json:"completed"  params:"type: int, nullable: false"`
	Priority  int           `json:"priority"   params:"type: int, nullable: false"`
	TaskGroup int           `json:"task_group" params:"type: int, nullable: false, foreign_key: task_group"`
	Name      string        `json:"name"       params:"type: varchar, length: 255, nullable: false"`
	Type      string        `json:"type"       params:"type: varchar, length: 255, nullable: false"`
	CustomId  string        `json:"custom_id"  params:"type: varchar, length: 255, nullable: false"`
}

func (t *Task) GetTaskGroup() *TaskGroup {
	return orm.FindById(TaskGroup{}, t.TaskGroup).(*TaskGroup)
}
