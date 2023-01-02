package main

import (
	"context"
	"fmt"

	surreal "github.com/garrison-henkle/surrealdb.go"
)

const (
	DB_URL      = "ws://localhost:8000/rpc"
	DB_USER     = "root"
	PORT = "8000"
	DB_PASSWORD = "root"
)


type DB = surreal.DB

func dbConn(ctx *context.Context) (*DB, error) {

  db, err := surreal.New(ctx, DB_URL)
	if err != nil {
		return nil, err
	}

	if err = db.Signin(ctx, surreal.UserInfo{
	    User: DB_USER,
	    Password: DB_PASSWORD,
	}); err != nil {
		return nil, fmt.Errorf("error authenticating to database")
	}

	return db, nil
}
