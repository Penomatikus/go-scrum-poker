package entity

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/penomatikus/go-scrum-poker/server/database"
)

type RoomToken string

type RoomRepository interface {
	Create(context.Context, *Room) error
	Update(context.Context, Room) error
	Delete(context.Context, RoomToken) error
}

type repoImpl struct {
	db *sqlx.DB
}

var _ RoomRepository = &repoImpl{}

func (repo *repoImpl) Create(ctx context.Context, room *Room) error {
	_, err := database.MustHaveTx(ctx).NamedExecContext(ctx, `
			INSERT INTO room (token, name, description) 
			VALUES (:token, :name, :description)
		`, room)
	return err
}

func (repo *repoImpl) Update(ctx context.Context, room Room) error {
	return nil
}

func (repo *repoImpl) Delete(ctx context.Context, token RoomToken) error {
	return nil
}

func ProvideRoomRepository(db *sqlx.DB) RoomRepository {
	repository := &repoImpl{db: db}
	repository.db.MustExec(Schema)
	return repository
}
