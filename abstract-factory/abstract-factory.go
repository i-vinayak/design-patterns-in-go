package main

import (
	"errors"
	"fmt"
)

type DatabaseFactory interface {
	NewSQLDatabase(sqlType string) (SqlDatabase, error)
	NewNoSQLDatabase(dbType string) (NoSqlDatabase, error)
}

type databaseFactory struct {
}

func (factory *databaseFactory) NewNoSQLDatabase(dbType string) (NoSqlDatabase, error) {
	switch dbType {
	case "mongo":
		return &mongoDb{data: make(map[string]map[string]string)}, nil
	case "cassandra":
		return &cassandraDb{data: map[string]map[string]string{}}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Failed to create NoSQL Database of type %s", dbType))
	}
}

func (factory *databaseFactory) NewSQLDatabase(sqlType string) (SqlDatabase, error) {
	switch sqlType {
	case "mysql":
		return &mySql{data: map[string]map[string]string{}}, nil
	case "postgres":
		return &postgres{data: map[string]map[string]string{}}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Failed to create Relational Database of type %s", sqlType))
	}
}

func NewDatabaseFactory() DatabaseFactory {
	return &databaseFactory{}
}
