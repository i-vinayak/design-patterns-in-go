package main

import "fmt"

type DatabaseProxy struct {
	db    Database
	cache map[string]string
}

func NewDatabaseService() *DatabaseProxy {
	return &DatabaseProxy{cache: make(map[string]string)}
}

// proxy to the actual database, checks if data is in cache before initializing the db
func (dp *DatabaseProxy) GetData(query string) string {
	data, exists := dp.cache[query]
	if exists {
		fmt.Println("Returned from Proxy")
		return data
	}

	dp.db = NewDatabase()
	data = dp.db.GetData(query)
	dp.cache[query] = data
	return data
}
