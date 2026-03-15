package models

type Car struct {
	ID          int     `json:"id"`
	Brand       string  `json:"brand"`
	Model       string  `json:"model"`
	Year        int     `json:"year"`
	PricePerDay float64 `json:"price_per_day"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"created_at"`
}