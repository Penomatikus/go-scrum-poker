package entity

import (
	"context"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/penomatikus/go-scrum-poker/server/database"
	"github.com/penomatikus/go-scrum-poker/server/database/txtest"
	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	ts := prepareRepoTestSetup(t)

	room := &Room{
		Token:       12345,
		Name:        "Test",
		Description: "This is a description",
	}

	txtest.AutoCommit(t, ts.db, func(ctx context.Context) error {
		return ts.repo.Create(ctx, room)
	})

	var roomDB Room
	assert.Nil(t, ts.db.Get(&roomDB, "SELECT * FROM room"))
	fmt.Printf("\n%#v\n", roomDB)
}

type repositoryTestSetup struct {
	db        *sqlx.DB
	ctx       context.Context
	txManager database.TransactionManger
	repo      RoomRepository
}

func prepareRepoTestSetup(t *testing.T) *repositoryTestSetup {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	txManager := database.ProvideTransactionManger(db)
	repo := ProvideRoomRepository(db)

	return &repositoryTestSetup{
		db:        db,
		ctx:       ctx,
		txManager: txManager,
		repo:      repo,
	}
}
