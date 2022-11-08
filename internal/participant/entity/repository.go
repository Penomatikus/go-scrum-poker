package entity

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type ParticipantRepository interface {
	Create(context.Context, *Participant) error
}

type repoImpl struct {
	db *sqlx.DB
}

var _ ParticipantRepository = &repoImpl{}

func (repo *repoImpl) Create(ctx context.Context, participant *Participant) error {
	tx := repo.db.MustBegin()
	tx.NamedExecContext(ctx, ``, participant)
	return tx.Commit()
}

func ProvideParticipantRpository(db *sqlx.DB) ParticipantRepository {
	repository := &repoImpl{db: db}
	repository.db.MustExec(Schema)
	return repository
}
