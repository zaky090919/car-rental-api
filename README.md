# Car Rental API

A RESTful API for managing a car rental system built with **Golang**, **Gorilla Mux**, and **PostgreSQL**.

This project allows users to manage cars, customers, and bookings for a car rental service.

---

# Tech Stack

* **Golang**
* **Gorilla Mux** (HTTP Router)
* **PostgreSQL**
* **SQL**

---

# Project Structure

```
car-rental-api
│
├── database
│   ├── schema.sql
│   └── seed.sql
│
├── handlers
│   ├── car_handler.go
│   ├── customer_handler.go
│   └── booking_handler.go
│
├── models
│   ├── car.go
│   ├── customer.go
│   └── booking.go
│
├── docs
│   └── erd.png
│
├── main.go
├── go.mod
└── .gitignore
```

---

# Installation

Clone the repository

```
git clone https://github.com/zaky090919/car-rental-api.git
cd car-rental-api
```

Install dependencies

```
go mod tidy
```

---

# Environment Variables

Create a `.env` file in the root directory:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=car_rental
DB_SSLMODE=disable
```

---

# Database Setup

Create database and run schema:

```
psql -U postgres -d car_rental -f database/schema.sql
```

Insert seed data:

```
psql -U postgres -d car_rental -f database/seed.sql
```

---

# Run the Application

```
go run main.go
```

Server will start at:

```
http://localhost:8080
```

---

# API Endpoints

## Cars

GET all cars

```
GET /cars
```

Create car

```
POST /cars
```

Update car

```
PUT /cars/{id}
```

Delete car

```
DELETE /cars/{id}
```

---

## Customers

```
GET /customers
POST /customers
PUT /customers/{id}
DELETE /customers/{id}
```

---

## Bookings

```
GET /bookings
POST /bookings
PUT /bookings/{id}
DELETE /bookings/{id}
```

---

# Database ERD

![ERD](docs/erd.png)

---

# Author

Zaky Al Fitra
