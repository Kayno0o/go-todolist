package entity

type User struct {
	Id            int    `json:"id"            params:"type: int, primary_key: true, auto_increment: true"`
	Name          string `json:"name"          params:"type: varchar, length: 255, nullable: false"`
	Email         string `json:"email"         params:"type: varchar, length: 255, nullable: false"`
	Password      string `json:"password"      params:"type: varchar, length: 255, nullable: false"`
	Discriminator string `json:"discriminator" params:"type: int,     length: 255, nullable: false"`
	Token         string `json:"token"         params:"type: varchar, length: 255, nullable: false"`
	Role          string `json:"role"          params:"type: varchar, length: 255, nullable: false"`
}
