package main

import (
	"fmt"

	surreal "github.com/surrealdb/surrealdb.go"
)

const (
	DB_URL      = "http://localhost:8000/rpc"
	DB_USER     = "root"
	DB_PASSWORD = "root"
)


type DB = surreal.DB

func dbConn() (*DB, error) {

	db, err := surreal.New(DB_URL)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to database")
	}

	_, err = db.Signin(map[string]interface{}{
		"user": DB_USER,
		"pass": DB_PASSWORD,
	})
	if err != nil {
		return nil, fmt.Errorf("Error authenticating to database")
	}

	return db, nil
}
