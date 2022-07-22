-- name: CreateUser :one
INSERT INTO users ( 
	id, hashed_password, full_name
) VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetEvent :one
SELECT * FROM events LIMIT 1 OFFSET $1;

-- name: CreateEvent :one
INSERT INTO events (
	"name", "start_date", "end_date", "link", "description"
) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetEventCount :one
SELECT COUNT(*) FROM events;