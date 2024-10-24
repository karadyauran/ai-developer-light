-- name: CreateUser :one
INSERT INTO users (github_id, username, email, token)
VALUES ($1, $2, $3, $4)
RETURNING id, github_id, username, email, token, created_at, updated_at;

-- name: GetUserByID :one
SELECT id, github_id, username, email, token, created_at, updated_at
FROM users
WHERE id = $1;

-- name: GetUserByGitHubID :one
SELECT id, github_id, username, email, token, created_at, updated_at
FROM users
WHERE github_id = $1;

-- name: UpdateUserToken :exec
UPDATE users
SET token = $2, updated_at = NOW()
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;