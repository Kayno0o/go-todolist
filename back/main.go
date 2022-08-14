package main

import (
	"Todolist/app/api"
	"Todolist/app/entity"
	"Todolist/app/orm"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	orm.RegisterEntity(&entity.User{})
	orm.RegisterEntity(&entity.UserIcon{})

	orm.RegisterEntity(&entity.TaskGroup{})
	orm.RegisterEntity(&entity.Task{})

	orm.RegisterEntity(&entity.UserTaskGroup{})

	api.StartApi()
}
