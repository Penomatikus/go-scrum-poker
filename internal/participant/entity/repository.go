package entity

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/penomatikus/go-scrum-poker/server/database"
)

type ParticipantRepository interface {
	Create(context.Context, *Participant) error
}

type repoImpl struct {
	db *sqlx.DB
}

var _ ParticipantRepository = &repoImpl{}

func (repo *repoImpl) Create(ctx context.Context, participant *Participant) error {
	_, err := database.MustHaveTx(ctx).NamedExecContext(ctx, ``, participant)
	return err
}

func ProvideParticipantRpository(db *sqlx.DB) ParticipantRepository {
	repository := &repoImpl{db: db}
	repository.db.MustExec(Schema)
	return repository
}
