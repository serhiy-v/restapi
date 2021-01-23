package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/serhiy-v/restapi/models"
	"net/http"
)

func GetProjects(w http.ResponseWriter, r *http.Request)  {
	json.NewEncoder(w).Encode(Projects)

}

func GetProject(w http.ResponseWriter, r *http.Request)  {
	param := mux.Vars(r)
	for _,item := range Projects {
		if item.ID == param["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Project{})

}

func CreateProject(w http.ResponseWriter, r *http.Request)  {

}

func UpdateProject(w http.ResponseWriter, r *http.Request)  {

}

func DeleteProject(w http.ResponseWriter, r *http.Request)  {

}