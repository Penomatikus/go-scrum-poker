package entity

import (
	"context"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	ts := prepareRepoTestSetup(t)

	room := &Room{
		Token:       12345,
		Name:        "Test",
		Description: "This is a description",
	}
	assert.Nil(t, ts.repo.Create(ts.ctx, room))

	var roomDB Room
	assert.Nil(t, ts.db.Get(&roomDB, "SELECT * FROM room"))
	fmt.Printf("\n%#v\n", roomDB)
}

type sqlxRepositoryTestSetup struct {
	db   *sqlx.DB
	ctx  context.Context
	repo RoomRepository
}

func prepareRepoTestSetup(t *testing.T) *sqlxRepositoryTestSetup {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	repo := ProvideRoomRepository(db)
	return &sqlxRepositoryTestSetup{
		db:   db,
		ctx:  ctx,
		repo: repo,
	}
}
