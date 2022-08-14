package orm

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
)

func GetStructName(i interface{}) string {
	return strcase.ToSnake(reflect.Indirect(reflect.ValueOf(i)).Type().Name())
}

func GetStructFields(i interface{}) []interface{} {
	var fields []interface{}

	vi := reflect.ValueOf(i)
	v := reflect.Indirect(vi)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fields = append(fields, field.Addr().Interface())
	}

	return fields
}

func GetStructFieldsByName(i interface{}) map[string]interface{} {
	var fields = make(map[string]interface{})

	vi := reflect.ValueOf(i)
	v := reflect.Indirect(vi)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fields[v.Type().Field(i).Name] = field.Addr().Interface()
	}

	return fields
}

func GetStructValuesByName(i interface{}) map[string]interface{} {
	var fields = make(map[string]interface{})

	vi := reflect.ValueOf(i)
	v := reflect.Indirect(vi)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fields[v.Type().Field(i).Name] = field.Interface()
	}

	return fields
}

func ToLower(str string) string {
	return strings.ToLower(str)
}

func FirstToUpper(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}

func ToJson(i interface{}) string {
	j, err := json.Marshal(i)

	if err != nil {
		return ""
	}

	return string(j)
}
