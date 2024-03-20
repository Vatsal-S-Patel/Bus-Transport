package controllers

import (
	"busproject/model"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Controller struct {
	DB *sql.DB
}

func NewController(db *sql.DB) *Controller {
	return &Controller{
		DB: db,
	}
}

func OutputToClient(w http.ResponseWriter, code int, message string, data any) {
	
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(model.OutputStruct{Code : code, Message: message,Data :data})
	
	if err != nil {
		log.Println("heyy can't send response back to client error ", err)
		log.Println("the inputs were ",code, message, data)
		w.Write([]byte("can't send response back if you are developer please refer log asap!!!"))
	}
	
}
