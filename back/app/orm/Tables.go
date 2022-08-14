package orm

import (
	"fmt"
)

func CreateTable(i interface{}) {
	tableName := GetStructName(i)
	params := Entities[tableName].Params

	query := "CREATE TABLE IF NOT EXISTS " + tableName + " ("

	index := 0
	for _, param := range params {
		query += param.Name + " " + param.Type

		if param.Length != -1 {
			query += "(" + fmt.Sprint(param.Length) + ")"
		}

		if param.PrimaryKey {
			query += " PRIMARY KEY"
		}

		if param.AutoIncrement {
			query += " AUTO_INCREMENT"
		}

		if param.Nullable {
			query += " NULL"
		} else {
			query += " NOT NULL"
		}

		if param.Default != "" {
			query += " DEFAULT " + param.Default
		}

		if param.Unique {
			query += " UNIQUE"
		}

		if param.ForeignKey != "" {
			query += ", FOREIGN KEY " + tableName + "_" + param.ForeignKey + " (" + param.Name + ") REFERENCES " + param.ForeignKey + "(id)"

			if param.OnDelete != "" {
				query += " ON DELETE " + param.OnDelete
			}
		}

		if index < len(params)-1 {
			query += ", "
		}
		index++
	}

	query += ")"

	_, err := Query(query)

	if err != nil {
		panic(err)
	}
}
