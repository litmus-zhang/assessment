-- name: CreateCompany :one
INSERT INTO company_details (name, address, phone_number, email, owned_by) VALUES ($1, $2, $3, $4, $5) RETURNING *;


-- name: GetCompany :one
SELECT * FROM company_details WHERE id = $1
LIMIT 1;

-- name: GetCompanyCreatedByUser :one
SELECT * FROM company_details WHERE owned_by = $1 AND id = $2 LIMIT 1;

-- name: GetCompaniesCreatedByUser :many
SELECT * FROM company_details WHERE owned_by = $1
LIMIT $2
OFFSET $3;


-- name: UpdateCompany :one
UPDATE company_details SET name = $1, address = $2, phone_number = $3, email = $4 WHERE id = $5 AND owned_by = $6 RETURNING *;

-- name: DeleteCompany :exec
DELETE FROM company_details WHERE id = $1 AND owned_by = $2 RETURNING *;