// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: categories.sql

package dao

import (
	"context"

	"github.com/google/uuid"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO categories(name, belongs_to, created_by, updated_by) 
VALUES (lower($1), $3, $2, $2)
RETURNING id, name, belongs_to, created_at, updated_at, deleted_at, created_by, updated_by
`

type CreateCategoryParams struct {
	Lower     string
	CreatedBy uuid.UUID
	BelongsTo uuid.UUID
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, createCategory, arg.Lower, arg.CreatedBy, arg.BelongsTo)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.BelongsTo,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
UPDATE categories
SET deleted_at = NOW()
WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, id)
	return err
}

const getCategoryById = `-- name: GetCategoryById :one
SELECT id, name, belongs_to, created_at, updated_at, deleted_at, created_by, updated_by FROM categories
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1
`

func (q *Queries) GetCategoryById(ctx context.Context, id int32) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategoryById, id)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.BelongsTo,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const getCategoryByName = `-- name: GetCategoryByName :one
SELECT id, name, belongs_to, created_at, updated_at, deleted_at, created_by, updated_by FROM categories
WHERE name = lower($1) AND deleted_at IS NULL
LIMIT 1
`

func (q *Queries) GetCategoryByName(ctx context.Context, lower string) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategoryByName, lower)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.BelongsTo,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const listCategories = `-- name: ListCategories :many
SELECT id, name, belongs_to, created_at, updated_at, deleted_at, created_by, updated_by FROM categories
WHERE deleted_at IS NULL
AND belongs_to = $1
ORDER BY name
`

func (q *Queries) ListCategories(ctx context.Context, belongsTo uuid.UUID) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, listCategories, belongsTo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.BelongsTo,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.CreatedBy,
			&i.UpdatedBy,
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

const updateCategory = `-- name: UpdateCategory :one
UPDATE categories
SET name = lower($2), 
updated_at = NOW(),
updated_by = $3
WHERE id = $1
RETURNING id, name, belongs_to, created_at, updated_at, deleted_at, created_by, updated_by
`

type UpdateCategoryParams struct {
	ID        int32
	Lower     string
	UpdatedBy uuid.UUID
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, updateCategory, arg.ID, arg.Lower, arg.UpdatedBy)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.BelongsTo,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}
