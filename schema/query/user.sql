-- name: CreateUser :one
INSERT INTO users (email, password, first_name, last_name)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1
LIMIT 1;

-- name: UpdateUser :one
UPDATE users SET email = $1, password = $2, first_name = $3, last_name = $4
WHERE id = $5 RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1 RETURNING *;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;