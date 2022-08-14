package orm

import (
	"strings"

	"github.com/iancoleman/strcase"
)

type EntityData struct {
	Entity     interface{}
	Params     []Param
	Visibility bool
}

var Entities = make(map[string]EntityData)

func RegisterEntity(i interface{}, visible ...bool) {
	isVisible := true

	if len(visible) > 0 {
		isVisible = visible[0]
	}

	Entities[GetStructName(i)] = EntityData{
		Entity:     i,
		Params:     GetStructParams(i),
		Visibility: isVisible,
	}
	CreateTable(i)
}

func All(i interface{}) []interface{} {
	return QueryEntities(i, "SELECT * FROM ${ð}")
}

func FindById(i interface{}, id int) interface{} {
	return QueryEntity(i, "SELECT * FROM ${ð} WHERE id = ?", id)
}

func FindByField(i interface{}, field string, value interface{}) interface{} {
	return QueryEntity(i, "SELECT * FROM ${ð} WHERE "+field+" LIKE ?", value)
}

func FindAllByField(i interface{}, field string, value interface{}) []interface{} {
	return QueryEntities(i, "SELECT * FROM ${ð} WHERE "+field+" LIKE ?", value)
}

func FindByFields(i interface{}, fields []string, values ...interface{}) interface{} {
	tableName := GetStructName(i)

	query := "SELECT * FROM " + tableName + " WHERE "
	where := []string{}

	for _, field := range fields {
		where = append(where, field+" = ?")
	}

	query += strings.Join(where, " AND ")

	return QueryEntity(i, query, values...)
}

func FindByUnique(i interface{}) interface{} {
	params := GetStructParams(i)
	tableName := GetStructName(i)

	structValues := GetStructValuesByName(i)

	query := "SELECT * FROM " + tableName + " WHERE "
	values := []interface{}{}
	where := []string{}

	for _, param := range params {
		if param.Unique {
			name := strcase.ToCamel(param.Name)
			where = append(where, name+" = ?")
			values = append(values, structValues[name])

			if structValues[name] == nil {
				return nil
			}
		}
	}

	query += strings.Join(where, " OR ")

	return QueryEntities(i, query, values...)[0]
}

func SaveEntity(i interface{}) (int, error) {
	fields := GetStructValuesByName(i)
	entity := FindById(i, fields["Id"].(int))

	if entity == nil {
		entity = FindByUnique(i)

		if entity == nil {
			return InsertEntity(i)
		} else {
			return UpdateEntity(i)
		}
	} else {
		return UpdateEntity(i)
	}
}

func InsertEntity(i interface{}) (int, error) {
	params := GetStructParams(i)
	tableName := GetStructName(i)

	fields := GetStructValuesByName(i)
	values := []interface{}{}

	insertStr := ""
	valuesStr := ""

	index := 0
	for _, param := range params {
		fieldName := strcase.ToCamel(param.Name)
		if fields[fieldName] == nil || fieldName == "Id" {
			continue
		}

		values = append(values, fields[fieldName])

		if index > 0 {
			insertStr += ", "
			valuesStr += ", "
		}

		insertStr += param.Name
		valuesStr += "?"

		index++
	}

	query := "INSERT INTO " + tableName + " (" + insertStr + ") VALUES (" + valuesStr + ") RETURNING id"

	rows, err := Query(query, values...)
	if err != nil {
		return 0, err
	}

	for rows.Next() {
		var id int
		rows.Scan(&id)
		fields["Id"] = id
	}

	return fields["Id"].(int), nil
}

func UpdateEntity(i interface{}) (int, error) {
	params := GetStructParams(i)
	tableName := GetStructName(i)

	fields := GetStructValuesByName(i)
	values := []interface{}{}

	updateStr := ""

	index := 0
	for _, param := range params {
		values = append(values, fields[param.Name])

		updateStr += param.Name + " = ?"

		if index < len(params)-1 {
			updateStr += ", "
		}
		index++
	}

	query := "UPDATE " + tableName + " SET " + updateStr + " WHERE id = ?"
	values = append(values, fields["Id"])

	Query(query, values...)

	return fields["Id"].(int), nil
}
