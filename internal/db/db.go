package db

import (
	"database/sql"
	"fmt"
	"hotel-system-1/internal/config"
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func InitDB(cfg config.Config) *sql.DB {
	var err error

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.SSLMode)

	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal("Cannot connect to the database: ", err)
	}
	fmt.Println("✅ Connected to hotel_booking database!")

	if err := goose.Up(dbConn, "internal/db/migrations"); err != nil {
		log.Fatalf("Error applying migrations: %v\n", err)
	}
	return dbConn
}
