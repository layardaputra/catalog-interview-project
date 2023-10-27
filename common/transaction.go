package common

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// Define a custom type for the context key
type contextKey string

const (
	// TransactionContextKey is the key for storing the transaction in the context
	TransactionContextKey contextKey = "transaction"
)

func GetTransactionFromContext(ctx context.Context, db *sqlx.DB) (*sqlx.Tx, error) {
	tx, ok := ctx.Value(TransactionContextKey).(*sqlx.Tx)
	if ok {
		return tx, nil // Transaction exists in the context
	}

	newTx, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	_ = context.WithValue(ctx, TransactionContextKey, newTx)

	return newTx, nil
}

// RunInTrans executes an operation within a transaction using a context.
func RunInTrans(ctx context.Context, db *sqlx.DB, fn func(ctx context.Context) error) error {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil || err != nil {
			_ = tx.Rollback()
		}
	}()

	// Add the transaction to the context
	ctx = context.WithValue(ctx, TransactionContextKey, tx)

	err = fn(ctx)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
