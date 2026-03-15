package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"car-rental-api/database"
	"car-rental-api/models"

	"github.com/gorilla/mux"
)

// GET /bookings
func GetBookings(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, customer_id, car_id, start_date, end_date, total_price, created_at FROM bookings")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var b models.Booking
		err := rows.Scan(&b.ID, &b.CustomerID, &b.CarID, &b.StartDate, &b.EndDate, &b.TotalPrice, &b.CreatedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bookings = append(bookings, b)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

// POST /bookings
func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var b models.Booking

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if b.CustomerID == 0 || b.CarID == 0 || b.StartDate == "" || b.EndDate == "" {
		http.Error(w, "customer_id, car_id, start_date, and end_date are required", http.StatusBadRequest)
		return
	}

	err = database.DB.QueryRow(
		"INSERT INTO bookings(customer_id, car_id, start_date, end_date, total_price) VALUES($1,$2,$3,$4,$5) RETURNING id",
		b.CustomerID, b.CarID, b.StartDate, b.EndDate, b.TotalPrice,
	).Scan(&b.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}

// PUT /bookings/{id}
func UpdateBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var b models.Booking
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if b.StartDate == "" || b.EndDate == "" {
		http.Error(w, "start_date and end_date must not be empty", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec(
		"UPDATE bookings SET customer_id=$1, car_id=$2, start_date=$3, end_date=$4, total_price=$5 WHERE id=$6",
		b.CustomerID, b.CarID, b.StartDate, b.EndDate, b.TotalPrice, id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Booking updated"})
}

// DELETE /bookings/{id}
func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	_, err := database.DB.Exec("DELETE FROM bookings WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Booking deleted"})
}
