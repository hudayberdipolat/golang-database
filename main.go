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
	//selectCity()
	// select one city
	//selectOneCity(1)
	selectWithPreparedStatemant("Ashgabat")
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
	rows, err := db.Query("select id, name , code FROM cities")

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

func selectOneCity(id int) {
	var city City
	sqlQuery := fmt.Sprintf("SELECT id, name, code FROM cities WHERE id = %v", id)

	err := db.QueryRow(sqlQuery).Scan(&city.ID, &city.Name, &city.Code)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(city)
	}
}

func selectWithPreparedStatemant(cityName string) {
	stmt, err := db.Prepare("SELECT id, name, code FROM cities WHERE name = $1")
	if err != nil {
		log.Fatal(err.Error())
		return
	} else {
		var city City
		err := stmt.QueryRow(cityName).Scan(&city.ID, &city.Name, &city.Code)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println(city)
		}
	}

}
