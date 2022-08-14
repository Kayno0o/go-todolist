package entity

type Job struct {
	Id   int    `json:"id"   params:"type: int, primary_key: true"`
	Name string `json:"name" params:"type: varchar, length: 255"`
}
