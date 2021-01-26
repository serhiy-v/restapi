package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/serhiy-v/restapi/db"
	"github.com/serhiy-v/restapi/models"
	"log"
	"net/http"
	"strconv"
)

func ShowTask(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	res,err := db.ShowTask(id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(res)
}

func CreateTask(w http.ResponseWriter, r *http.Request)  {
	var p models.Task
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	err = db.CreateTask(p)
	if err != nil{
		log.Fatal(err)
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request)  {
	var p models.Task
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"])
	if err != nil {
		log.Fatal(err)
	}
	err = db.UpdateTask(p,id)
	if err != nil{
		log.Fatal(err)
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request)  {
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"])
	if err != nil {
		log.Fatal(err)
	}
	err = db.DeleteTask(id)
	if err != nil{
		log.Fatal(err)
	}
}
