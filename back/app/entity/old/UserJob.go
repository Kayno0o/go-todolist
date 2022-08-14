package entity

import "Todolist/app/orm"

type UserJob struct {
	Id      int  `json:"id"       params:"type: int, primary_key: true, auto_increment: true"`
	User    int  `json:"user"     params:"type: int, foreign_key: user"`
	Job     int  `json:"job"      params:"type: int, foreign_key: job"`
	Level   int  `json:"level"    params:"type: int, default: 0"`
	Visible bool `json:"visible"  params:"type: tinyint, default: 0"`
}

func (u *UserJob) GetUser() *User {
	return orm.FindById(&User{}, u.User).(*User)
}

func (u *UserJob) GetJob() *Job {
	return orm.FindById(&Job{}, u.Job).(*Job)
}
