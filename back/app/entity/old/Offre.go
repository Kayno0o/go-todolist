package entity

type Offre struct {
	Id    int    `json:"id"    params:"type: int, primary_key: true, auto_increment: true"`
	Price int    `json:"price" params:"type: int, default: 0"`
	Name  string `json:"name"  params:"type: varchar, length: 255"`
}
