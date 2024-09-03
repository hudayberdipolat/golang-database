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

type City struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code int    `json:"code"`
}

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
	//insertCity()
	// select city
	selectCity()
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

func selectCity() {
	var cityList []City
	rows, err := db.Query("select id, name , code from cities")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		for rows.Next() {
			var city City
			if errScan := rows.Scan(&city.ID, &city.Name, &city.Code); errScan != nil {
				log.Println(errScan.Error())
			} else {
				cityList = append(cityList, city)
			}
		}
		rows.Close()
		fmt.Println(cityList)
	}
}
