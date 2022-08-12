-- name: CreateUser :one
INSERT INTO users ( 
	id, hashed_password, full_name
) VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;
