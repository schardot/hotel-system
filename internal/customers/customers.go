package customers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Customers struct {
	dbConnection *sql.DB
}

func NewCustomers(dbConnection *sql.DB) Customers {
	return Customers{dbConnection: dbConnection}
}
func (c *Customers) ShowCustomerForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/customerform.html")
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func (c *Customers) SubmitCustomerForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")
	telNum := r.FormValue("tel_number")
	celNum := r.FormValue("cel_number")
	idType := r.FormValue("id_type")
	idNumber := r.FormValue("id_number")
	address := r.FormValue("address")
	zipCode := r.FormValue("zip_code")
	city := r.FormValue("city")
	state := r.FormValue("state")
	country := r.FormValue("country")

	_, err := c.dbConnection.Exec("INSERT INTO customers (first_name, last_name,  email, tel_number, cel_number, id_type, id_number, address, zip_code, city, state, country) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
		firstName, lastName, email, telNum, celNum, idType, idNumber, address, zipCode, city, state, country)

	if err != nil {
		log.Printf("❌ Database error: %v\n", err)
		http.Error(w, "Error saving booking", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Thank you, %s! Your booking has been received.", firstName)
}
