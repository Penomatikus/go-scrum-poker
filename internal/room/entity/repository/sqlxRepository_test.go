package repository

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/penomatikus/go-scrum-poker/internal/room/entity"
	"github.com/penomatikus/go-scrum-poker/server/database"
	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	ts := prepareRepoTestSetup(t)

	room := &entity.Room{
		Token:       "12345",
		Name:        "Test",
		Description: "This is a description",
	}
	assert.Nil(t, ts.sqlxRepository.Create(ts.ctx, room))

	var roomDB entity.Room
	assert.Nil(t, ts.db.Get(&roomDB, "SELECT * FROM room"))
}

type sqlxRepositoryTestSetup struct {
	db             *sqlx.DB
	ctx            context.Context
	sqlxRepository SqlxRepository
}

func prepareRepoTestSetup(t *testing.T) *sqlxRepositoryTestSetup {
	orm, err := database.ProvideORM(database.SQLX, database.DialectSqlite3, ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	repo := ProvideSqlxRepository(orm)
	return &sqlxRepositoryTestSetup{
		db:             repo.db,
		ctx:            ctx,
		sqlxRepository: repo,
	}
}
