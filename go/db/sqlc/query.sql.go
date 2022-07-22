// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: query.sql

package db

import (
	"context"
	"time"
)

const createEvent = `-- name: CreateEvent :one
INSERT INTO events (
	"name", "start_date", "end_date", "link", "description"
) VALUES ($1, $2, $3, $4, $5) RETURNING name, start_date, end_date, link, description
`

type CreateEventParams struct {
	Name        string
	StartDate   time.Time
	EndDate     time.Time
	Link        string
	Description string
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error) {
	row := q.db.QueryRowContext(ctx, createEvent,
		arg.Name,
		arg.StartDate,
		arg.EndDate,
		arg.Link,
		arg.Description,
	)
	var i Event
	err := row.Scan(
		&i.Name,
		&i.StartDate,
		&i.EndDate,
		&i.Link,
		&i.Description,
	)
	return i, err
}

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

const getEvent = `-- name: GetEvent :one
SELECT name, start_date, end_date, link, description FROM events LIMIT 1 OFFSET $1
`

func (q *Queries) GetEvent(ctx context.Context, offset int32) (Event, error) {
	row := q.db.QueryRowContext(ctx, getEvent, offset)
	var i Event
	err := row.Scan(
		&i.Name,
		&i.StartDate,
		&i.EndDate,
		&i.Link,
		&i.Description,
	)
	return i, err
}

const getEventCount = `-- name: GetEventCount :one
SELECT COUNT(*) FROM events
`

func (q *Queries) GetEventCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getEventCount)
	var count int64
	err := row.Scan(&count)
	return count, err
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
