package entity

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	roomEntity "github.com/penomatikus/go-scrum-poker/internal/room/entity"
	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	ts := prepareRepoTestSetup(t)
	ts.createRoomInDB(t, 1)

	ts.participantRepo.Create(ts.ctx, &Participant{
		Name:      "Theo Test",
		RoomToken: 1,
	})

}

type repositoryTestSetup struct {
	db              *sqlx.DB
	ctx             context.Context
	roomRepo        roomEntity.RoomRepository
	participantRepo ParticipantRepository
}

func prepareRepoTestSetup(t *testing.T) *repositoryTestSetup {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	participantRepo := ProvideParticipantRpository(db)
	roomRepo := roomEntity.ProvideRoomRepository(db)
	return &repositoryTestSetup{
		db:              db,
		ctx:             ctx,
		roomRepo:        roomRepo,
		participantRepo: participantRepo,
	}
}

func (ts *repositoryTestSetup) createRoomInDB(t *testing.T, token int64) {
	assert.Nil(t, ts.roomRepo.Create(ts.ctx, &roomEntity.Room{
		Token:       token,
		Name:        "Test",
		Description: "This is a description",
	}))
}
