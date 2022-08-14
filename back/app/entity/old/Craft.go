package entity

import (
	"Todolist/app/orm"
	"database/sql"
)

type Craft struct {
	Id     int           `json:"id"     params:"type: int, primary_key: true"`
	Result int           `json:"result" params:"type: int, foreign_key: item"`
	Price  int           `json:"price"  params:"type: int"`
	Job    sql.NullInt16 `json:"job"    params:"type: int, foreign_key: job, nullable: true"`
}

func (c *Craft) GetResult() *Item {
	return orm.FindById(&Item{}, c.Result).(*Item)
}

func (c *Craft) GetJob() *Job {
	return orm.FindById(&Job{}, int(c.Job.Int16)).(*Job)
}
