package entity

type ConfigItem struct {
	Id    int    `json:"id"      params:"type: int, primary_key: true, auto_increment: true"`
	Name  string `json:"name"    params:"type: varchar, length: 255, nullable: false"`
	Value string `json:"value"   params:"type: text, nullable: false"`
}
