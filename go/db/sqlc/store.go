package db

import (
	"database/sql"
)

// >> defines all functions to execute db queries and transactions
// Created for the API
type Store interface {
	Querier
	// add transfers here
}

// >> provides function implementation to execute SQL queries
// >> also provides the transaction execution object
type SQLStore struct {
	db *sql.DB
	*Queries
}

// >> creates a new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// >> executes a function within a database transaction
// func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
// 	tx, err := store.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	q := New(tx)
// 	err = fn(q)
// 	if err != nil {
// 		if rbErr := tx.Rollback(); rbErr != nil {
// 			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
// 		}
// 		return err
// 	}

// 	return tx.Commit()
// }
