package handlers

import (
	"github.com/gorilla/mux"
	"github.com/serhiy-v/restapi/models"
)

var Projects []models.Project

func NewRouter() *mux.Router {
	Projects = append(Projects, models.Project{ID: "1",Name: "prj", Description: "simply prj"})
	r := mux.NewRouter()
    //Project Route
	r.HandleFunc("/projects",GetProjects).Methods("GET")
	r.HandleFunc("/projects/{id}",GetProject).Methods("GET")
	r.HandleFunc("/projects",CreateProject).Methods("POST")
	r.HandleFunc("/projects/{id}",UpdateProject).Methods("PUT")
	r.HandleFunc("/projects/{id}",DeleteProject).Methods("DELETE")

	r.HandleFunc("/columns",ShowColumn).Methods("GET")
	r.HandleFunc("/columns/{id}",UpdateColumn).Methods("PUT")
	r.HandleFunc("/columns/{id}",DeleteColumn).Methods("DELETE")

	r.HandleFunc("/tasks",ShowTask).Methods("GET")
	r.HandleFunc("/tasks/{id}",UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}",DeleteTask).Methods("DELETE")

	r.HandleFunc("/comments",ShowComment).Methods("GET")
	r.HandleFunc("/comments/{id}",UpdateComment).Methods("PUT")
	r.HandleFunc("/comments/{id}",DeleteComment).Methods("DELETE")

	return r
}
