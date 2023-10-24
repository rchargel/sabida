-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = lower($1)
AND deleted_at IS NULL
LIMIT 1;

-- name: GetUsersByEmail :one
SELECT * FROM users
WHERE email = lower($1)
AND deleted_at IS NULL
LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
WHERE deleted_at IS NULL
ORDER BY username;

-- name: CreateUser :one
INSERT INTO users (username, email)
VALUES (lower($1), lower($2))
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET username = $2,
	email = $3,
	password = $4,
	update_at = now()
WHERE id = $1
RETURNING *;

-- name: UpdatePassword :exec
UPDATE users
SET password = $2
WHERE id = $1;

-- name: DeleteUser :exec
UPDATE users SET deleted_at = NOW()
WHERE id = $1;

-- name: GetOrganizationByID :one
SELECT * FROM organizations
WHERE id = $1
LIMIT 1;

-- name: GetOrganizationByName :one
SELECT * FROM organizations
WHERE name = lower($1)
LIMIT 1;

-- name: GetOrganizationsByUser :many
SELECT organizations.* FROM organizations INNER JOIN
user_organizations ON organizations.id = user_organizations.organization_id
AND user_organizations.deleted_at IS NULL
WHERE user_organizations.user_id = $1
AND organizations.deleted_at IS NULL;

-- name: GetUsersByOrganization :many
SELECT users.* FROM users INNER JOIN
user_organizations ON users.id = user_organizations.user_id
AND user_organizations.deleted_at IS NULL
WHERE user_organizations.organization_id = $1
AND users.deleted_at IS NULL;

-- name: AddUserToOrganization :exec
INSERT INTO user_organizations (user_id, organization_id)
VALUES ($1, $2);

-- name: RemoveUserFromOrganization :exec
UPDATE user_organizations
SET deleted_at = NOW(), deleted_by = $3
WHERE user_id = $1 AND organization_id = $2;