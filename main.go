package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"hotel-system-1/internal/config"
	"hotel-system-1/internal/customers"
	"hotel-system-1/internal/db"

	"github.com/sethvargo/go-envconfig"
)

func main() {
	var cfg config.Config
	ctx := context.Background()
	if err := envconfig.Process(ctx, &cfg); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	dbConnection := db.InitDB(cfg)
	defer dbConnection.Close()

	customer := customers.NewCustomers(dbConnection)
	http.HandleFunc("/customerform", customer.ShowCustomerForm)
	http.HandleFunc("/submitcustomer", customer.SubmitCustomerForm)
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
