package graph

import (
	surreal "github.com/garrison-henkle/surrealdb.go"
)
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
  Db *surreal.DB
}
