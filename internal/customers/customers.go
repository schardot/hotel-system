package customers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"hotel-system-1/internal/entities"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type CustomersService struct {
	dbConnection *sql.DB
}

func (c *CustomersService) GetAllCustomers() ([]entities.Customer, error) {
	rows, err := c.dbConnection.Query("SELECT * FROM customers")
	if err != nil {
		log.Println("Error querying customers:", err)
		return nil, err
	}
	defer rows.Close()

	var customers []entities.Customer

	for rows.Next() {
		var customer entities.Customer
		if err := rows.Scan(
			&customer.ID,
			&customer.FirstName,
			&customer.LastName,
			&customer.Email,
			&customer.TelNum,
			&customer.CelNum,
			&customer.IdType,
			&customer.IdNum,
			&customer.Address,
			&customer.ZipCode,
			&customer.City,
			&customer.State,
			&customer.Country,
			&customer.CreatedAt,
		); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (c *CustomersService) HandleGetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := c.GetAllCustomers()
	if err != nil {
		http.Error(w, "Error retrieving customers", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(customers); err != nil {
		http.Error(w, "Error encoding customers to JSON", http.StatusInternalServerError)
	}
}

func NewCustomers(dbConnection *sql.DB) CustomersService {
	return CustomersService{dbConnection: dbConnection}
}

func (c *CustomersService) ListCustomers(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/customers.html")
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func (c *CustomersService) ShowCustomerForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/customerform.html")
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func (c *CustomersService) SubmitCustomerForm(w http.ResponseWriter, r *http.Request) {
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
