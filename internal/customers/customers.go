package customers

import (
	"database/sql"
	"fmt"
	"hotel-system-1/internal/entities"
	"html/template"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

type CustomersService struct {
	dbConnection *sql.DB
}

func (c *CustomersService) GetAllCustomers(query string, args []interface{}) ([]entities.Customer, error) {
	rows, err := c.dbConnection.Query(query, args...)
	if err != nil {
		log.Println("Error querying customers:", err)
		return nil, err
	}
	defer rows.Close()

	var customers []entities.Customer

	for rows.Next() {
		var customer entities.Customer
		if err := rows.Scan(
			&customer.CustomerID,
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
	searchQuery := r.URL.Query().Get("search")
	query := "SELECT * FROM customers"

	var args []interface{}
	var conditions []string

	if searchQuery != "" {
		conditions = append(conditions, "(CAST(customer_id AS TEXT) ILIKE $1 OR first_name ILIKE $1 OR last_name ILIKE $1 OR email ILIKE $1)")
		formattedSearch := "%" + searchQuery + "%"
		args = append(args, formattedSearch)
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	customers, err := c.GetAllCustomers(query, args)
	if err != nil {
		http.Error(w, "Error retrieving customers", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	tmpl, err := template.ParseFiles("templates/customers.html")
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, customers); err != nil {
		log.Printf("Error rendering template: %v", err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func NewCustomers(dbConnection *sql.DB) CustomersService {
	return CustomersService{dbConnection: dbConnection}
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

	customerID := r.FormValue("customer_id")
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

	if customerID == "" {
		_, err := c.dbConnection.Exec("INSERT INTO customers (first_name, last_name,  email, tel_number, cel_number, id_type, id_number, address, zip_code, city, state, country) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
			firstName, lastName, email, telNum, celNum, idType, idNumber, address, zipCode, city, state, country)
		if err != nil {
			log.Printf("❌ Database error: %v\n", err)
			http.Error(w, "Error saving booking", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "✅ New customer created: %s %s", firstName, lastName)
	} else {
		_, err := c.dbConnection.Exec("UPDATE customers SET first_name=$1, last_name=$2, email=$3, tel_number=$4, cel_number=$5, id_type=$6, id_number=$7, address=$8, zip_code=$9, city=$10, state=$11, country=$12 WHERE customer_id=$13", firstName, lastName, email, telNum, celNum, idType, idNumber, address, zipCode, city, state, country, customerID)
		if err != nil {
			log.Printf("❌ Database error: %v\n", err)
			http.Error(w, "Error updating customer", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "✅ Customer updated: %s %s", firstName, lastName)
	}
}

func (c *CustomersService) EditCustomerForm(w http.ResponseWriter, r *http.Request) {
	customerID := r.URL.Query().Get("customer_id")
	if customerID == "" {
		http.Error(w, "Missing customer ID", http.StatusBadRequest)
		return
	}

	var cstmr entities.Customer
	err := c.dbConnection.QueryRow("SELECT * FROM customers WHERE customer_id = $1", customerID).Scan(
		&cstmr.CustomerID,
		&cstmr.FirstName,
		&cstmr.LastName,
		&cstmr.Email,
		&cstmr.TelNum,
		&cstmr.CelNum,
		&cstmr.IdType,
		&cstmr.IdNum,
		&cstmr.Address,
		&cstmr.ZipCode,
		&cstmr.City,
		&cstmr.State,
		&cstmr.Country,
		&cstmr.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Customer not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Error querying customer by customer id", http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles("templates/customerform.html")
	if err != nil {
		http.Error(w, "Error loading form template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, cstmr); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
