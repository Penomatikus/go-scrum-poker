package adapter

import (
	"database/sql"

	"github.com/penomatikus/go-scrum-poker/internal/room/entity"
	"github.com/penomatikus/go-scrum-poker/server/database"
)

var _ entity.RoomRepository = &SqlxRepositoryAdapter{}
var _ database.Adapter = &SqlxRepositoryAdapter{}

type SqlxRepositoryAdapter struct {
	database.Adapter
	db *sql.DB
}

func (sqlxA *SqlxRepositoryAdapter) Create(request entity.RoomCreateRequest) (*entity.Room, error) {
	return nil, nil
}

func (sqlxA *SqlxRepositoryAdapter) Update(request entity.RoomUpdateRequest) error {
	return nil
}

func (sqlxA *SqlxRepositoryAdapter) Delete(token entity.RoomToken) error {
	return nil
}

func (sqlxA *SqlxRepositoryAdapter) migrateSchema() error {
	return nil
}

func ProvideSqlxRepositoryAdapter(db *sql.DB) entity.RoomRepository {
	repositoryAdapter := &SqlxRepositoryAdapter{db: db}
	if err := repositoryAdapter.migrateSchema(); err != nil {
		panic("Error!")
	}
	return repositoryAdapter
}
