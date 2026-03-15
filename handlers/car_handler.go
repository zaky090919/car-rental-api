package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"car-rental-api/database"
	"car-rental-api/models"

	"github.com/gorilla/mux"
)

func GetCars(w http.ResponseWriter, r *http.Request) {

	rows, err := database.DB.Query("SELECT id, brand, model, year, price_per_day, status, created_at FROM cars")

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer rows.Close()

	var cars []models.Car

	for rows.Next() {

		var c models.Car

		rows.Scan(&c.ID, &c.Brand, &c.Model, &c.Year, &c.PricePerDay, &c.Status, &c.CreatedAt)

		cars = append(cars, c)
	}

	json.NewEncoder(w).Encode(cars)
}

func CreateCar(w http.ResponseWriter, r *http.Request) {

	var c models.Car

	json.NewDecoder(r.Body).Decode(&c)

	err := database.DB.QueryRow(
		"INSERT INTO cars(brand, model, year, price_per_day, status) VALUES($1,$2,$3,$4,$5) RETURNING id",
		c.Brand, c.Model, c.Year, c.PricePerDay, c.Status,
	).Scan(&c.ID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(c)
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var c models.Car

	json.NewDecoder(r.Body).Decode(&c)

	_, err := database.DB.Exec(
		"UPDATE cars SET brand=$1, model=$2, year=$3, price_per_day=$4, status=$5 WHERE id=$6",
		c.Brand, c.Model, c.Year, c.PricePerDay, c.Status, id,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode("Car updated")
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	_, err := database.DB.Exec("DELETE FROM cars WHERE id=$1", id)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode("Car deleted")
}