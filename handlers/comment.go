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

func ShowComment(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	res,err := db.ShowComment(id)
	json.NewEncoder(w).Encode(res)
}

func CreateComment(w http.ResponseWriter, r *http.Request)  {
	var p models.Comment
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	err = db.CreateComment(p)
	if err != nil{
		log.Fatal(err)
	}
}

func UpdateComment(w http.ResponseWriter, r *http.Request)  {
	var p models.Comment
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
	err = db.UpdateComment(p,id)
	if err != nil{
		log.Fatal(err)
	}
}

func DeleteComment(w http.ResponseWriter, r *http.Request)  {
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"])
	if err != nil {
		log.Fatal(err)
	}
	err = db.DeleteComment(id)
	if err != nil{
		log.Fatal(err)
	}
}
