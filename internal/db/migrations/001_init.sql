-- +goose Up
CREATE TABLE IF NOT EXISTS customers (
    customer_id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    tel_number VARCHAR(20) NOT NULL,
    cel_number VARCHAR(20) NOT NULL,
    id_type TEXT NOT NULL,
    id_number VARCHAR(20) NOT NULL,
    address TEXT NOT NULL,
    zip_code VARCHAR(10) NOT NULL,
    city TEXT NOT NULL,
    state TEXT NOT NULL,
    country TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE customers;