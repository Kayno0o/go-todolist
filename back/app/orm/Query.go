package orm

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

func QueryEntity(i interface{}, query string, args ...interface{}) interface{} {
	entities := QueryEntities(i, query, args...)

	if len(entities) > 0 {
		return entities[0]
	}

	return nil
}

func QueryEntities(i interface{}, query string, args ...interface{}) []interface{} {
	query = strings.Replace(query, "${ð}", GetStructName(i), 1)
	rows, err := Query(query, args...)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var instances []interface{}

	for rows.Next() {
		instance := reflect.New(reflect.Indirect(reflect.ValueOf(i)).Type()).Interface()

		err := rows.Scan(GetStructFields(instance)...)

		if err != nil {
			fmt.Println(err)
			return nil
		}

		instances = append(instances, instance)
	}

	return instances
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	db, _ := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PSWD")+"@/"+os.Getenv("DB_NAME"))

	res, err := db.Query(query, args...)

	log.Println(query)
	log.Println(args...)

	if err != nil {
		fmt.Println(query)
		fmt.Println(args...)
		fmt.Println(err)
	}

	defer db.Close()

	return res, err
}
