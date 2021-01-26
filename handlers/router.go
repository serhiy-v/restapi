package handlers

import (
	"github.com/gorilla/mux"
)


func NewRouter() *mux.Router {
	r := mux.NewRouter()
    //Project Route
	r.HandleFunc("/projects",GetProjects).Methods("GET")
	r.HandleFunc("/projects/{id}",GetProject).Methods("GET")
	r.HandleFunc("/projects",CreateProject).Methods("POST")
	r.HandleFunc("/projects/{id}",UpdateProject).Methods("PUT")
	r.HandleFunc("/projects/{id}",DeleteProject).Methods("DELETE")
	//Column Route
	r.HandleFunc("/columns/{id}",ShowColumn).Methods("GET")
	r.HandleFunc("/columns",CreateColumn).Methods("POST")
	r.HandleFunc("/columns/{id}",UpdateColumn).Methods("PUT")
	r.HandleFunc("/columns/{id}",DeleteColumn).Methods("DELETE")
	//Task Route
	r.HandleFunc("/tasks/{id}",ShowTask).Methods("GET")
	r.HandleFunc("/tasks",CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}",UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}",DeleteTask).Methods("DELETE")
	//Comment Route
	r.HandleFunc("/comments/{id}",ShowComment).Methods("GET")
	r.HandleFunc("/comments",CreateComment).Methods("POST")
	r.HandleFunc("/comments/{id}",UpdateComment).Methods("PUT")
	r.HandleFunc("/comments/{id}",DeleteComment).Methods("DELETE")

	return r
}
