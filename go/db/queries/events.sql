
-- name: GetEvent :one
SELECT * FROM events LIMIT 1 OFFSET $1;

-- name: CreateEvent :one
INSERT INTO events (
	"name", "start_date", "end_date", "link", "description"
) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetEventCount :one
SELECT COUNT(*) FROM events;