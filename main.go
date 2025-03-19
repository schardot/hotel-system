package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

var db *sql.DB

func bookPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func submitBooking(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.FormValue("name")
	email := r.FormValue("email")
	checkInStr := r.FormValue("check_in")
	checkOutStr := r.FormValue("check_out")
	guests := r.FormValue("guests")

	// Convert the date strings to time.Time
	checkIn, err := time.Parse("2006-01-02", checkInStr) // Date format must match the input
	if err != nil {
		log.Printf("Invalid check-in date format: %v\n", err)
		http.Error(w, "Invalid check-in date", http.StatusBadRequest)
		return
	}

	checkOut, err := time.Parse("2006-01-02", checkOutStr)
	if err != nil {
		log.Printf("Invalid check-out date format: %v\n", err)
		http.Error(w, "Invalid check-out date", http.StatusBadRequest)
		return
	}

	// Convert guests to integer
	guestsInt, err := strconv.Atoi(guests)
	if err != nil {
		log.Printf("Invalid number of guests: %v\n", err)
		http.Error(w, "Invalid number of guests", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO bookings (name, email, check_in, check_out, guests) VALUES ($1, $2, $3, $4, $5)",
		name, email, checkIn, checkOut, guestsInt)

	if err != nil {
		log.Printf("❌ Database error: %v\n", err)
		http.Error(w, "Error saving booking", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Thank you, %s! Your booking has been received.", name)
}

func main() {
	db = connectDB()
	defer db.Close()

	http.HandleFunc("/book", bookPage)
	http.HandleFunc("/submit", submitBooking)
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
