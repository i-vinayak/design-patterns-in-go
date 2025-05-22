package main

import (
	"errors"
	"fmt"
)

type Database interface {
	PutData(string)
	GetData() string
}

// mongodb struct implements Database interface
type MongoDB struct {
	currentData string
}

func (mongoDB *MongoDB) GetData() string {
	return mongoDB.currentData + " (from mongodb)"
}
func (mongoDB *MongoDB) PutData(data string) {
	mongoDB.currentData = data
}

// sql struct implements Database interface
type SQL struct {
	currentData string
}

func (sql *SQL) GetData() string {
	return sql.currentData + " (from sql)"
}
func (sql *SQL) PutData(data string) {
	sql.currentData = data
}

// This is the database factory - but follows go idiomatic naming convention
func NewDatabase(dbType string) (Database, error) {
	switch dbType {
	case "relational":
		return &SQL{}, nil
	case "nosql":
		return &MongoDB{}, nil
	default:
		fmt.Printf("Database of type %s is not supported\n", dbType)
		return nil, errors.New("Failed to created database - invalid dbType")
	}
}
