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

func ShowColumn(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	res,err := db.ShowCol(id)
	json.NewEncoder(w).Encode(res)
}

func CreateColumn(w http.ResponseWriter, r *http.Request)  {
	var p models.Column
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	err = db.CreateCol(p)
	if err != nil{
		log.Fatal(err)
	}
}

func UpdateColumn(w http.ResponseWriter, r *http.Request)  {
	var p models.Column
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
	err = db.UpdateCol(p,id)
	if err != nil{
		log.Fatal(err)
	}

}

func DeleteColumn(w http.ResponseWriter, r *http.Request)  {
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"])
	if err != nil {
		log.Fatal(err)
	}
	err = db.DeleteCol(id)
	if err != nil{
		log.Fatal(err)
	}
}
