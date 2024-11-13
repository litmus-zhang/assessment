-- name: CreateInvoice :one
INSERT INTO invoices ( customer_id ,name, due_date, status, company_id, note, discount ) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: UpdateInvoice :one
UPDATE invoices SET name = $2, due_date = $3, status = $4, note = $5, discount = $6 WHERE id = $1 AND company_id =$7 RETURNING *;

-- name: GetOneInvoice :one
SELECT * FROM invoices WHERE id = $1 AND company_id =$2
LIMIT 1;

-- name: GetAllInvoices :many
SELECT * FROM invoices WHERE company_id = $1 ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetInvoicesByStatus :many
SELECT * FROM invoices WHERE status = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3;




-- name: GetInvoiceTotalFromItems :exec
UPDATE invoices
SET amount = (
    SELECT 
        CAST(SUM(total_price) AS numeric)
    FROM
        items
    WHERE
        invoice_id = $1
)
WHERE id = $1;


-- name: GetCompanyInvoiceSummary :many
SELECT 
    status,
    COUNT(*) AS count,
    CAST(SUM(amount) AS numeric) AS total_amount
FROM 
    invoices
WHERE 
    company_id = $1
GROUP BY 
    status;


-- name: DeleteInvoice :exec
DELETE FROM invoices WHERE id = $1 AND company_id =$2;