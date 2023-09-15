-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: UpdateUserNameById :one
UPDATE users SET name = $2, updated_at = $3 WHERE id = $1
RETURNING *;

-- name: DeleteUserById :exec
DELETE FROM users WHERE id = $1;

-- name: UserExistsById :one
SELECT EXISTS(SELECT 1 FROM users WHERE id = $1) AS exists;