-- name: CreatePaymentDetail :one
INSERT INTO payment_details (account_name, account_number, bank_name, company_id) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetACompanyPaymentDetailByID :one
SELECT * FROM payment_details WHERE id = $1 AND company_id = $2
LIMIT 1;

-- name: ListAllCompanyPaymentDetails :many
SELECT * FROM payment_details WHERE company_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdatePaymentDetail :one
UPDATE payment_details SET account_name = $1, account_number = $2, bank_name = $3 WHERE id = $4 AND company_id = $5 RETURNING *;

-- name: DeletePaymentDetail :exec
DELETE FROM payment_details WHERE id = $1 AND company_id = $2 RETURNING *;
