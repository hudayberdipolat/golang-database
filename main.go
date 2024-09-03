package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var (
	db    *sql.DB
	dbErr error
)

func main() {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "12345"
	dbName := "golang-database"
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s, password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	db, dbErr = sql.Open("postgres", psqlInfo)
	defer db.Close()
	if dbErr != nil {
		panic(dbErr)
	}
	fmt.Println("Connection database successfully")
}

/* Note
C - create -> insert
R - read   -> select
U - update -> update
D - delete -> delete
*/
