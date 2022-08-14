package api

import (
	"Todolist/app/orm"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func StartApi() {
	router := httprouter.New()

	router.GET("/api/:object/all", GetAllObject)
	router.GET("/api/:object/all/:by/:value", GetAllObjectBy)
	router.GET("/api/:object/one/:by/:value", GetOneObjectBy)

	router.POST("/api/:object/new", CreateObject)

	port := os.Getenv("API_PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Println("API started on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func GetOneObjectBy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	objectName, by, value := ps.ByName("object"), ps.ByName("by"), ps.ByName("value")

	entity := GetEntityFromName(w, objectName)
	if entity == nil {
		SendError(w, "Entity not found", 500)
		return
	}

	object := orm.FindByField(entity, by, value)
	if object == nil {
		SendJsonResponse(w, []byte("[]"))
		return
	}

	jsonObject, err := json.Marshal(object)
	if err != nil {
		SendError(w, "Error while marshalling object", 500)
		return
	}

	SendJsonResponse(w, jsonObject)
}

func GetAllObject(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	objectName := ps.ByName("object")

	entity := GetEntityFromName(w, objectName)
	if entity == nil {
		SendError(w, "Entity not found", 500)
		return
	}

	objects := orm.All(entity)
	if objects == nil {
		SendJsonResponse(w, []byte("[]"))
		return
	}

	jsonObjects, err := json.Marshal(objects)
	if err != nil {
		SendError(w, "Error while marshalling object", 500)
		return
	}

	SendJsonResponse(w, jsonObjects)
}

func GetAllObjectBy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	objectName, by, value := ps.ByName("object"), ps.ByName("by"), ps.ByName("value")

	entity := GetEntityFromName(w, objectName)
	if entity == nil {
		SendError(w, "Entity not found", 500)
		return
	}

	objects := orm.FindAllByField(entity, by, value)
	if objects == nil {
		SendJsonResponse(w, []byte("[]"))
		return
	}

	jsonObjects, err := json.Marshal(objects)
	if err != nil {
		SendError(w, "Error while marshalling object", 500)
		return
	}

	SendJsonResponse(w, jsonObjects)
}

func CreateObject(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	objectName := ps.ByName("object")

	entity := GetEntityFromName(w, objectName)
	if entity == nil {
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil && data == nil {
		SendError(w, "Error while reading body", 500)
		return
	}

	json.Unmarshal(data, entity)

	id, err := orm.InsertEntity(entity)
	if err != nil {
		SendError(w, err.Error(), 500)
		return
	}

	if id == 0 {
		SendError(w, "Error while inserting entity", 500)
		return
	}

	SendJsonResponse(w, []byte("{\"id\": \""+strconv.Itoa(id)+"\"}"))
}

func SendError(w http.ResponseWriter, message string, errorCode int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("{\"status\": \"" + strconv.Itoa(errorCode) + "\", \"error\": \"" + message + "\"}"))
}

func SendJsonResponse(w http.ResponseWriter, jsonObject []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonObject))
}

func GetEntityFromName(w http.ResponseWriter, name string) interface{} {
	object := orm.Entities[name]

	if object.Entity == nil || !object.Visibility {
		SendJsonResponse(w, []byte("[]"))
		return nil
	}

	return object.Entity
}
