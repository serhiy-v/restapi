package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() error {
	var err error
	db, err = sql.Open("postgres","user=postgres password=root dbname=restapi sslmode=disable")
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil{
		return err
	}
	return nil
}