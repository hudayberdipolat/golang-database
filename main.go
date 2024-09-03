package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var (
	db    *sql.DB
	dbErr error
)

func main() {
	host := "127.0.0.1"
	port := "5432"
	user := "postgres"
	password := "12345"
	dbName := "golang-database"
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		user, password, dbName, host, port)
	// Open the connection
	db, dbErr = sql.Open("postgres", connStr)
	if dbErr != nil {
		panic(fmt.Sprintf("Failed to open a DB connection: %v", dbErr))
	}
	defer db.Close()
	if dbErr != nil {
		panic(dbErr)
	}
	err := db.Ping()
	if err != nil {
		log.Fatal("error database ping ", err.Error())
	}

	fmt.Println("Connection database successfully")

	// call insert function
	insertCity()
}

/* Note
C - create -> insert
R - read   -> select
U - update -> update
D - delete -> delete
*/

func insertCity() {
	result, err := db.Exec("INSERT INTO cities(name, code) values('Ashgabat', 744000)")
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println(result.RowsAffected())
	}
}
