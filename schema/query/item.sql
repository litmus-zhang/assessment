-- name: CreateItem :one
INSERT INTO items ( name, unit_price, description, invoice_id, quantity) 
VALUES ($1, $2, $3, $4, $5) RETURNING *;


-- name: UpdateItem :one
UPDATE items SET name = $2, unit_price = $3, quantity=$4  WHERE id = $1 AND invoice_id =$5 RETURNING *;

-- name: DeleteItem :exec
DELETE FROM items WHERE id = $1;



-- name: GetAlltemsForAnInvoice :many
SELECT * FROM items WHERE invoice_id = $1
LIMIT $2 OFFSET $3;



