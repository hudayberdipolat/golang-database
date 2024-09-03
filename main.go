package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/hudayberdipolat/golang-database/entity"
	"github.com/hudayberdipolat/golang-database/repository"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

	cityRepo := repository.NewCityRepo(db)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "unsupported http method", http.StatusMethodNotAllowed)
			return
		}
		cityList := cityRepo.FindAllCities()
		json.NewEncoder(w).Encode(cityList)
	})

	http.HandleFunc("/city", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			if r.URL.Query().Has("id") {
				id, _ := strconv.Atoi(r.URL.Query().Get("id"))
				city := cityRepo.GetByID(id)
				if city == nil {
					w.WriteHeader(http.StatusNotFound)
					json.NewEncoder(w).Encode("Not Found")
					return
				}
				w.WriteHeader(http.StatusOK)
				if err := json.NewEncoder(w).Encode(city); err != nil {
					http.Error(w, err.Error(), http.StatusMethodNotAllowed)
					return
				}
			}
			if r.URL.Query().Has("name") {
				name := r.URL.Query().Get("name")
				city := cityRepo.GetByName(name)
				if city == nil {
					w.WriteHeader(http.StatusNotFound)
					json.NewEncoder(w).Encode("Not Found")
					return
				}
				w.WriteHeader(http.StatusOK)
				if err := json.NewEncoder(w).Encode(city); err != nil {
					http.Error(w, err.Error(), http.StatusMethodNotAllowed)
					return
				}
			}
			return
		}
		http.Error(w, "unsupported http method", http.StatusMethodNotAllowed)
		return
	})

	http.HandleFunc("/city/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "unsupported http method", http.StatusMethodNotAllowed)
			return
		}
		var city entity.City

		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		cityRepo.InsertCity(city)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(city)
	})

	go func() {
		errServer := http.ListenAndServe(":8080", nil)
		if errServer != nil {
			fmt.Println("Error starting HTTP server", errServer)
		}

	}()
	fmt.Println("HTTP server started")
	<-make(chan struct{})

	// call insert function
	//insertCity()
	// select city
	//selectCity()
	// select one city
	//selectOneCity(1)
	//selectWithPreparedStatemant("Ashgabat")
}

/* Note
C - create -> insert
R - read   -> select
U - update -> update
D - delete -> delete
*/
//
//func insertCity() {
//	result, err := db.Exec("INSERT INTO cities(name, code) values('Ashgabat', 744000)")
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	} else {
//		fmt.Println(result.RowsAffected())
//	}
//}
//
//func selectCity() {
//	var cityList []City
//	rows, err := db.Query("select id, name , code FROM cities")
//
//	if err != nil {
//		fmt.Println(err.Error())
//	} else {
//		for rows.Next() {
//			var city City
//			if errScan := rows.Scan(&city.ID, &city.Name, &city.Code); errScan != nil {
//				log.Println(errScan.Error())
//			} else {
//				cityList = append(cityList, city)
//			}
//		}
//		rows.Close()
//		fmt.Println(cityList)
//	}
//}
//
//func selectOneCity(id int) {
//	var city City
//	sqlQuery := fmt.Sprintf("SELECT id, name, code FROM cities WHERE id = %v", id)
//
//	err := db.QueryRow(sqlQuery).Scan(&city.ID, &city.Name, &city.Code)
//	if err != nil {
//		fmt.Println(err.Error())
//	} else {
//		fmt.Println(city)
//	}
//}
//
//func selectWithPreparedStatemant(cityName string) {
//	stmt, err := db.Prepare("SELECT id, name, code FROM cities WHERE name = $1")
//	if err != nil {
//		log.Fatal(err.Error())
//		return
//	} else {
//		var city City
//		err := stmt.QueryRow(cityName).Scan(&city.ID, &city.Name, &city.Code)
//		if err != nil {
//			fmt.Println(err)
//			return
//		} else {
//			fmt.Println(city)
//		}
//	}
//}
