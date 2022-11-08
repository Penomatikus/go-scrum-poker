package database

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type TxFn func(ctx context.Context) error

type TransactionManger interface {
	WithTransaction(context.Context, TxFn) error
}

func ProvideTransactionManger(db *sqlx.DB) TransactionManger {
	return &tmImpl{db: db}
}

type tmImpl struct {
	db *sqlx.DB
}

var _ TransactionManger = &tmImpl{}

func (tm *tmImpl) WithTransaction(ctx context.Context, fn TxFn) (err error) {
	tx, ok := ctx.Value(TransaktionCtxKey).(*sqlx.Tx)
	if !ok {
		tx, ctx = beginTransaction(ctx, tm.db)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(ctx)
	return
}

type transaction string

const TransaktionCtxKey transaction = "TX_Key"

func beginTransaction(ctx context.Context, db *sqlx.DB) (*sqlx.Tx, context.Context) {
	tx := db.MustBeginTx(ctx, nil)
	return tx, context.WithValue(ctx, TransaktionCtxKey, tx)
}
