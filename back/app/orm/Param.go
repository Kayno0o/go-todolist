package orm

import (
	"reflect"
	"strconv"
	"strings"
)

type Param struct {
	Name          string
	Type          string
	Length        int
	Nullable      bool
	PrimaryKey    bool
	AutoIncrement bool
	ForeignKey    string
	OnDelete      string
	OnUpdate      string
	Default       string
	Unique        bool
	Uniques       map[string]string
}

func NewParam(param string) Param {
	var res Param = Param{"id", "int", -1, false, false, false, "", "", "", "", false, make(map[string]string)}

	params := strings.Split(param, ", ")

	for _, v := range params {
		if strings.Contains(v, ":") {
			param := strings.Split(v, ": ")
			switch param[0] {
			case "name":
				res.Name = param[1]
			case "type":
				res.Type = param[1]
			case "primary_key":
				res.PrimaryKey = param[1] == "true"
			case "auto_increment":
				res.AutoIncrement = param[1] == "true"
			case "nullable":
				res.Nullable = param[1] == "true"
			case "foreign_key":
				res.ForeignKey = param[1]
			case "on_delete":
				res.OnDelete = param[1]
			case "on_update":
				res.OnUpdate = param[1]
			case "length":
				res.Length, _ = strconv.Atoi(param[1])
			case "default":
				res.Default = param[1]
			case "unique":
				res.Unique = param[1] == "true"
			}
		}
	}

	return res
}

func GetStringParam(i interface{}, param string) string {
	vi := reflect.ValueOf(i)
	v := reflect.Indirect(vi)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)

		if field.Tag.Get("json") == NewParam(param).Name {
			paramsStr := field.Tag.Get("params")
			jsonStr := field.Tag.Get("json")

			return "name: " + jsonStr + ", " + paramsStr
		}
	}

	return ""
}

func GetStringParams(i interface{}) []string {
	vi := reflect.ValueOf(i)
	v := reflect.Indirect(vi)
	t := v.Type()

	params := make([]string, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)

		paramsStr := field.Tag.Get("params")
		jsonStr := field.Tag.Get("json")

		param := "name: " + jsonStr + ", " + paramsStr
		params[i] = param
	}

	return params
}

func GetStructParam(i interface{}, param string) Param {
	return NewParam(GetStringParam(i, param))
}

func GetStructParams(i interface{}) []Param {
	params := []Param{}

	for _, k := range GetStringParams(i) {
		params = append(params, GetStructParam(i, k))
	}

	return params
}
