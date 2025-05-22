package main

import "fmt"

func main() {
	factory := NewDatabaseFactory()
	fmt.Println("Got new database factory")
	fmt.Println(factory)
	fmt.Println("================")

	sqlDb, _ := factory.NewSQLDatabase("mysql")
	noSqlDb, _ := factory.NewNoSQLDatabase("mongo")

	fmt.Println("sqlDB: ", sqlDb)
	fmt.Println("Nosql DB: ", noSqlDb)
}
