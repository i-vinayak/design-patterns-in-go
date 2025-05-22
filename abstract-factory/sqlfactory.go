package main

type SqlDatabase interface {
	GetRow(id string) map[string]string
	PutRow(id string, row map[string]string) string
}

type mySql struct {
	data map[string]map[string]string
}

func (db *mySql) GetRow(id string) map[string]string {
	return db.data[id]
}
func (db *mySql) PutRow(id string, row map[string]string) string {
	db.data[id] = row
	return "Added to mysql database"
}

type postgres struct {
	data map[string]map[string]string
}

func (db *postgres) GetRow(id string) map[string]string {
	return db.data[id]
}
func (db *postgres) PutRow(id string, row map[string]string) string {
	db.data[id] = row
	return "Added to Postgres database"
}
