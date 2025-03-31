package bookings

import (
	"database/sql"
	"net/http"
	"text/template"
)

type BookingsService struct {
	dbConnection *sql.DB
}

func NewBookings(dbConnection *sql.DB) BookingsService {
	return BookingsService{dbConnection: dbConnection}
}

func (b *BookingsService) HandleBookingsPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/bookings.html")
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func (b *BookingsService) HandleNewBookingForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/bookingform.html")
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
