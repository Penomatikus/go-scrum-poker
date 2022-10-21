package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/penomatikus/go-scrum-poker/internal/room/entity"
	"github.com/penomatikus/go-scrum-poker/server/database"
)

var _ entity.RoomRepository = &SqlxRepository{}

type SqlxRepository struct {
	db *sqlx.DB
}

func (repo *SqlxRepository) Create(ctx context.Context, room *entity.Room) error {
	tx := repo.db.MustBegin()
	tx.NamedExecContext(ctx, "INSERT INTO room (token, name, description) VALUES (:token, :name, :description)", room)
	// repo.db.MustExec("INSERT INTO room (token, name, description) VALUES (?, ?, ?)", room.Token, room.Name, room.Description)
	return tx.Commit()
}

func (repo *SqlxRepository) Update(ctx context.Context, room entity.Room) error {
	return nil
}

func (repo *SqlxRepository) Delete(ctx context.Context, token entity.RoomToken) error {
	return nil
}

func ProvideSqlxRepository(orm *database.ORM) SqlxRepository {
	repository := SqlxRepository{db: orm.DB.(*sqlx.DB)}
	repository.db.MustExec(entity.Schema)
	return repository
}
