package entity

type User struct {
	Id       int    `json:"id"          params:"type: int, primary_key: true"`
	IsActive bool   `json:"is_active"   params:"type: tinyint"`
	Discord  string `json:"discord"     params:"type: varchar, length: 255, length: 255"`
	Dofus    string `json:"dofus"       params:"type: varchar, length: 255, length: 255"`
	Access   string `json:"access"      params:"type: datetime"`
}
