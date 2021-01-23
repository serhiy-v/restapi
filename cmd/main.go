package main

import (
	"github.com/serhiy-v/restapi/db"
	"github.com/serhiy-v/restapi/handlers"
	"log"
	"net/http"
)

func main(){
	err := db.Connect()
	if err != nil{
		log.Fatal(err)
	}
	s := &http.Server{
		Addr: ":8080",
		Handler: handlers.NewRouter(),
	}
	log.Fatal(s.ListenAndServe())
}
