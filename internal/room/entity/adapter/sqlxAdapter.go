package adapter

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/penomatikus/go-scrum-poker/internal/room/entity"
	"github.com/penomatikus/go-scrum-poker/server/database"
)

var _ entity.RoomRepository = &SqlxRepositoryAdapter{}
var _ database.Adapter = &SqlxRepositoryAdapter{}

type SqlxRepositoryAdapter struct {
	database.Adapter
	db *sqlx.DB
}

func (repo *SqlxRepositoryAdapter) Create(request entity.RoomCreateRequest) (*entity.Room, error) {
	return nil, nil
}

func (repo *SqlxRepositoryAdapter) Update(request entity.RoomUpdateRequest) error {
	return nil
}

func (repo *SqlxRepositoryAdapter) Delete(token entity.RoomToken) error {
	return nil
}

func (adapter *SqlxRepositoryAdapter) migrateSchema() error {
	return nil
}

func (adapter *SqlxRepositoryAdapter) mustGetDB(ctx context.Context) *sqlx.DB {
	if db := ctx.Value(database.DatabaseCtxKey).(*sqlx.DB); db != nil {
		return db
	}
	panic(fmt.Errorf("no DB instance was found during transaktion"))
}

func ProvideSqlxRepositoryAdapter(orm *database.ORM) entity.RoomRepository {
	repositoryAdapter := &SqlxRepositoryAdapter{db: orm.DB.(*sqlx.DB)}
	if err := repositoryAdapter.migrateSchema(); err != nil {
		panic("Error!")
	}
	return repositoryAdapter
}
