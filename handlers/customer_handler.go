package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"car-rental-api/database"
	"car-rental-api/models"

	"github.com/gorilla/mux"
)

func GetCustomers(w http.ResponseWriter, r *http.Request) {

	rows, err := database.DB.Query("SELECT id, name, nik, phone_number, created_at FROM customers")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer rows.Close()

	var customers []models.Customer

	for rows.Next() {

		var c models.Customer

		rows.Scan(&c.ID, &c.Name, &c.NIK, &c.PhoneNumber, &c.CreatedAt)

		customers = append(customers, c)
	}

	json.NewEncoder(w).Encode(customers)
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {

	var c models.Customer

	json.NewDecoder(r.Body).Decode(&c)

	err := database.DB.QueryRow(
		"INSERT INTO customers(name, nik, phone_number) VALUES($1,$2,$3) RETURNING id",
		c.Name, c.NIK, c.PhoneNumber,
	).Scan(&c.ID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(c)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var c models.Customer

	json.NewDecoder(r.Body).Decode(&c)

	_, err := database.DB.Exec(
		"UPDATE customers SET name=$1, nik=$2, phone_number=$3 WHERE id=$4",
		c.Name, c.NIK, c.PhoneNumber, id,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode("Customer updated")
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	_, err := database.DB.Exec("DELETE FROM customers WHERE id=$1", id)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode("Customer deleted")
}