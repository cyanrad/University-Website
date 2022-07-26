// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users ( 
	id, hashed_password, full_name
) VALUES ($1, $2, $3)
RETURNING id, hashed_password, full_name, created_at
`

type CreateUserParams struct {
	ID             string
	HashedPassword string
	FullName       string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.ID, arg.HashedPassword, arg.FullName)
	var i User
	err := row.Scan(
		&i.ID,
		&i.HashedPassword,
		&i.FullName,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, hashed_password, full_name, created_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.HashedPassword,
		&i.FullName,
		&i.CreatedAt,
	)
	return i, err
}
