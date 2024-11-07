// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: company.sql

package db

import (
	"context"
)

const createCompany = `-- name: CreateCompany :one
INSERT INTO company_details (name, address, phone_number, email, owned_by) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, address, phone_number, email, owned_by, created_at
`

type CreateCompanyParams struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	OwnedBy     int32  `json:"owned_by"`
}

func (q *Queries) CreateCompany(ctx context.Context, arg CreateCompanyParams) (CompanyDetail, error) {
	row := q.db.QueryRowContext(ctx, createCompany,
		arg.Name,
		arg.Address,
		arg.PhoneNumber,
		arg.Email,
		arg.OwnedBy,
	)
	var i CompanyDetail
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Address,
		&i.PhoneNumber,
		&i.Email,
		&i.OwnedBy,
		&i.CreatedAt,
	)
	return i, err
}

const deleteCompany = `-- name: DeleteCompany :exec
DELETE FROM company_details WHERE id = $1 AND owned_by = $2 RETURNING id, name, address, phone_number, email, owned_by, created_at
`

type DeleteCompanyParams struct {
	ID      int64 `json:"id"`
	OwnedBy int32 `json:"owned_by"`
}

func (q *Queries) DeleteCompany(ctx context.Context, arg DeleteCompanyParams) error {
	_, err := q.db.ExecContext(ctx, deleteCompany, arg.ID, arg.OwnedBy)
	return err
}

const getCompaniesCreatedByUser = `-- name: GetCompaniesCreatedByUser :many
SELECT id, name, address, phone_number, email, owned_by, created_at FROM company_details WHERE owned_by = $1
LIMIT $2
OFFSET $3
`

type GetCompaniesCreatedByUserParams struct {
	OwnedBy int32 `json:"owned_by"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) GetCompaniesCreatedByUser(ctx context.Context, arg GetCompaniesCreatedByUserParams) ([]CompanyDetail, error) {
	rows, err := q.db.QueryContext(ctx, getCompaniesCreatedByUser, arg.OwnedBy, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CompanyDetail{}
	for rows.Next() {
		var i CompanyDetail
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Address,
			&i.PhoneNumber,
			&i.Email,
			&i.OwnedBy,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCompany = `-- name: GetCompany :one
SELECT id, name, address, phone_number, email, owned_by, created_at FROM company_details WHERE id = $1
LIMIT 1
`

func (q *Queries) GetCompany(ctx context.Context, id int64) (CompanyDetail, error) {
	row := q.db.QueryRowContext(ctx, getCompany, id)
	var i CompanyDetail
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Address,
		&i.PhoneNumber,
		&i.Email,
		&i.OwnedBy,
		&i.CreatedAt,
	)
	return i, err
}

const getCompanyCreatedByUser = `-- name: GetCompanyCreatedByUser :one
SELECT id, name, address, phone_number, email, owned_by, created_at FROM company_details WHERE owned_by = $1 AND id = $2 LIMIT 1
`

type GetCompanyCreatedByUserParams struct {
	OwnedBy int32 `json:"owned_by"`
	ID      int64 `json:"id"`
}

func (q *Queries) GetCompanyCreatedByUser(ctx context.Context, arg GetCompanyCreatedByUserParams) (CompanyDetail, error) {
	row := q.db.QueryRowContext(ctx, getCompanyCreatedByUser, arg.OwnedBy, arg.ID)
	var i CompanyDetail
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Address,
		&i.PhoneNumber,
		&i.Email,
		&i.OwnedBy,
		&i.CreatedAt,
	)
	return i, err
}

const updateCompany = `-- name: UpdateCompany :one
UPDATE company_details SET name = $1, address = $2, phone_number = $3, email = $4 WHERE id = $5 AND owned_by = $6 RETURNING id, name, address, phone_number, email, owned_by, created_at
`

type UpdateCompanyParams struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	ID          int64  `json:"id"`
	OwnedBy     int32  `json:"owned_by"`
}

func (q *Queries) UpdateCompany(ctx context.Context, arg UpdateCompanyParams) (CompanyDetail, error) {
	row := q.db.QueryRowContext(ctx, updateCompany,
		arg.Name,
		arg.Address,
		arg.PhoneNumber,
		arg.Email,
		arg.ID,
		arg.OwnedBy,
	)
	var i CompanyDetail
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Address,
		&i.PhoneNumber,
		&i.Email,
		&i.OwnedBy,
		&i.CreatedAt,
	)
	return i, err
}