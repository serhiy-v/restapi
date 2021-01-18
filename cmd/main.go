package main

import (
	"log"
	"net/http"
	"github.com/serhiy-v/restapi/handlers"
)

func main(){
	s := &http.Server{
		Addr: ":8080",
		Handler: handlers.NewRouter(),
	}
	log.Fatal(s.ListenAndServe())
}
