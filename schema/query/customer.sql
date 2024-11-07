-- name: CreateCustomer :one
INSERT INTO customers (first_name, last_name, email, phone_number, company_id, address) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetCustomerByID :one
SELECT * FROM customers WHERE id = $1 AND company_id = $2
LIMIT 1;


-- name: GetCustomerByEmail :one
SELECT * FROM customers WHERE email = $1 AND company_id = $2
LIMIT 1;

-- name: UpdateCustomer :one
UPDATE customers SET first_name = $1, last_name = $2, email = $3, phone_number = $4, address = $7 WHERE id = $5 AND company_id = $6 RETURNING *;

-- name: DeleteCustomer :exec
DELETE FROM customers WHERE id = $1 AND company_id = $2 RETURNING *;

-- name: ListCustomers :many
SELECT * FROM customers WHERE company_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;