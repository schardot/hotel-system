-- +goose Up
ALTER TABLE customers DROP CONSTRAINT customers_email_key;

-- +goose Down

ALTER TABLE customers ADD CONSTRAINT customers_email_key UNIQUE (email);