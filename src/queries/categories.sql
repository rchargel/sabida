-- name: ListCategories :many
SELECT * FROM categories
WHERE deleted_at IS NULL
AND belongs_to = $1
ORDER BY name;

-- name: GetCategoryById :one
SELECT * FROM categories
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: GetCategoryByName :one
SELECT * FROM categories
WHERE name = lower($1) AND deleted_at IS NULL
LIMIT 1;

-- name: UpdateCategory :one
UPDATE categories
SET name = lower($2), 
updated_at = NOW(),
updated_by = $3
WHERE id = $1
RETURNING *;

-- name: DeleteCategory :exec
UPDATE categories
SET deleted_at = NOW()
WHERE id = $1;

-- name: CreateCategory :one
INSERT INTO categories(name, belongs_to, created_by, updated_by) 
VALUES (lower($1), $3, $2, $2)
RETURNING *;

