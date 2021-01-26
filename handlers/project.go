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

func GetProjects(w http.ResponseWriter, r *http.Request)  {
	res, err := db.Projects()
	if err != nil {
		 log.Fatal(err)
	}
	for _, prj := range res {
		json.NewEncoder(w).Encode(prj)
	}

}

func GetProject(w http.ResponseWriter, r *http.Request)  {
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"])
	if err != nil {
		log.Fatal(err)
	}
	res,err := db.Project(id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(res)

}

func CreateProject(w http.ResponseWriter, r *http.Request)  {
	var p models.Project
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	err = db.Insert(p)
	if err != nil{
		log.Fatal(err)
	}
}

func UpdateProject(w http.ResponseWriter, r *http.Request)  {
	var p models.Project
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
	err = db.Update(p,id)
	if err != nil{
		log.Fatal(err)
	}

}

func DeleteProject(w http.ResponseWriter, r *http.Request)  {
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"])
	if err != nil {
		log.Fatal(err)
	}
	err = db.Delete(id)
	if err != nil{
		log.Fatal(err)
	}
}