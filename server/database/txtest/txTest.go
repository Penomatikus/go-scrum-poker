package txtest

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/penomatikus/go-scrum-poker/server/database"
)

// AutoCommit is a bare bones implementation of a transaction without handling panics and not performing rollbacks.
// It takes testsetupDB and provides a commited transaction for txfn.
func AutoCommit(t *testing.T, testsetupDB *sqlx.DB, txfn database.TxFn) {
	if t == nil {
		panic("You shall not pass. AutoCommit is only allowed for testing!")
	}

	ctx := context.Background()
	tx := testsetupDB.MustBeginTx(ctx, nil)

	if err := txfn(context.WithValue(ctx, database.TransaktionCtxKey, tx)); err != nil {
		t.Fatal(err)
	}
	tx.Commit()
}
