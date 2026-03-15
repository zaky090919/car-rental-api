package routes

import (
	"car-rental-api/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {

	r := mux.NewRouter()

	// Customers
	r.HandleFunc("/customers", handlers.GetCustomers).Methods("GET")
	r.HandleFunc("/customers", handlers.CreateCustomer).Methods("POST")
	r.HandleFunc("/customers/{id}", handlers.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{id}", handlers.DeleteCustomer).Methods("DELETE")

	// Cars
	r.HandleFunc("/cars", handlers.GetCars).Methods("GET")
	r.HandleFunc("/cars", handlers.CreateCar).Methods("POST")
	r.HandleFunc("/cars/{id}", handlers.UpdateCar).Methods("PUT")
	r.HandleFunc("/cars/{id}", handlers.DeleteCar).Methods("DELETE")

	// Bookings
	r.HandleFunc("/bookings", handlers.GetBookings).Methods("GET")
	r.HandleFunc("/bookings", handlers.CreateBooking).Methods("POST")
	r.HandleFunc("/bookings/{id}", handlers.UpdateBooking).Methods("PUT")
	r.HandleFunc("/bookings/{id}", handlers.DeleteBooking).Methods("DELETE")

	return r
}