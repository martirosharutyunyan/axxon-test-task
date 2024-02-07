-- name: Create :one
INSERT INTO tasks (url)
VALUES ($1)
RETURNING *;

-- name: Update :exec
UPDATE tasks SET headers = $2, status = $3, length = $4 WHERE id = $1;


-- name: GetById :one
SELECT id, created_at, updated_at, url, headers, length, status
FROM tasks WHERE id = $1;