package main

type NoSqlDatabase interface {
	GetRow(id string) map[string]string
	PutRow(id string, row map[string]string) string
}

type mongoDb struct {
	data map[string]map[string]string
}

func (db *mongoDb) GetRow(id string) map[string]string {
	return db.data[id]
}
func (db *mongoDb) PutRow(id string, row map[string]string) string {
	db.data[id] = row
	return "Added to mysql database"
}

type cassandraDb struct {
	data map[string]map[string]string
}

func (db *cassandraDb) GetRow(id string) map[string]string {
	return db.data[id]
}
func (db *cassandraDb) PutRow(id string, row map[string]string) string {
	db.data[id] = row
	return "Added to Postgres database"
}
