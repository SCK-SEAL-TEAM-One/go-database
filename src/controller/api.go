package controller

import (
	"encoding/json"
	"model"
	"net/http"
	"service"
)

type Connect struct {
	DB string
	Id string
}

func Read(responseWriter http.ResponseWriter, request *http.Request) {

	// db := request.URL.Query().Get("db")
	// id := request.URL.Query().Get("id")

	// connectData := &Connect{
	// 	DB: db,
	// 	Id: id}

	var sayhi model.Sayhi = service.ReadSayhiFormLevelDB()

	JSONConnectData, _ := json.Marshal(sayhi)

	responseWriter.Write(JSONConnectData)
}

func Add(responseWriter http.ResponseWriter, request *http.Request) {

	db := request.URL.Query().Get("db")
	id := request.URL.Query().Get("id")

	connectData := &Connect{
		DB: db,
		Id: id}
	JSONConnectData, _ := json.Marshal(connectData)

	responseWriter.Write(JSONConnectData)
}

func Edit(responseWriter http.ResponseWriter, request *http.Request) {

	db := request.URL.Query().Get("db")
	id := request.URL.Query().Get("id")

	connectData := &Connect{
		DB: db,
		Id: id}
	JSONConnectData, _ := json.Marshal(connectData)

	responseWriter.Write(JSONConnectData)
}

func Delete(responseWriter http.ResponseWriter, request *http.Request) {

	db := request.URL.Query().Get("db")
	id := request.URL.Query().Get("id")

	connectData := &Connect{
		DB: db,
		Id: id}
	JSONConnectData, _ := json.Marshal(connectData)

	responseWriter.Write(JSONConnectData)
}
