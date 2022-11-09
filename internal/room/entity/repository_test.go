package entity

import (
	"context"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
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
	ts.db.Get(&roomDB, "SELECT * FROM room")
	assert.Equal(t, room, &roomDB)
	fmt.Printf("\n%#v\n", roomDB)
}

type repositoryTestSetup struct {
	db   *sqlx.DB
	repo RoomRepository
}

func prepareRepoTestSetup(t *testing.T) *repositoryTestSetup {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	return &repositoryTestSetup{
		db:   db,
		repo: ProvideRoomRepository(db),
	}
}
