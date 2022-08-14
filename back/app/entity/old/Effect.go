package entity

type Effect struct {
	Id          int    `json:"id"          params:"type: int, primary_key: true"`
	Description string `json:"description" params:"type: varchar, length: 255"`
}
