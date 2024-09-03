package repository

import (
	"database/sql"
	"fmt"
	"github.com/hudayberdipolat/golang-database/entity"
)

type CityRepo struct {
	db *sql.DB
}

func NewCityRepo(db *sql.DB) *CityRepo {
	return &CityRepo{db}
}

func (c *CityRepo) InsertCity(city entity.City) {
	stmt, err := c.db.Prepare("INSERT INTO cities(name, code) VALUES ($1, $2)")
	result, err := stmt.Exec(city.Name, city.Code)
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println(result.RowsAffected())
	}
}

func (c *CityRepo) FindAllCities() []entity.City {
	var cities []entity.City
	rows, err := c.db.Query("SELECT id, name, code FROM cities")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	} else {
		for rows.Next() {
			var city entity.City
			err := rows.Scan(&city.ID, &city.Name, &city.Code)
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}
			cities = append(cities, city)
		}
		if err := rows.Close(); err != nil {
			fmt.Println(" error rows close ", err.Error())
			return nil
		}
	}
	return cities
}

func (c CityRepo) GetByID(id int) *entity.City {
	var city entity.City
	sql := "SELECT id, name, code FROM cities WHERE id = $1"
	stmt, err := c.db.Prepare(sql)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	} else {
		err := stmt.QueryRow(id).Scan(&city.ID, &city.Name, &city.Code)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
	}
	fmt.Println(city)
	return &city
}

func (c CityRepo) GetByName(name string) *entity.City {
	var city entity.City
	sql := "SELECT id, name, code FROM cities WHERE name = $1"
	stmt, err := c.db.Prepare(sql)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	} else {
		err := stmt.QueryRow(name).Scan(&city.ID, &city.Name, &city.Code)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
	}
	fmt.Println(city)
	return &city
}
