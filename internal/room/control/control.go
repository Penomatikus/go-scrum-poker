package control

import (
	"context"

	"github.com/penomatikus/go-scrum-poker/internal/room/entity"
	"github.com/penomatikus/go-scrum-poker/server/database"
)

type (
	RoomRequest struct {
		Name        string
		Description string
	}
)

type (
	RoomController interface {
		Create(context.Context, RoomRequest) (*entity.Room, error)
		Update(context.Context, RoomRequest) error
		Delete(context.Context, string) error
	}

	roomControllerImpl struct {
		txManager database.TransactionManger
		repo      entity.RoomRepository
	}
)

var _ RoomController = &roomControllerImpl{}

func (c *roomControllerImpl) Create(ctx context.Context, request RoomRequest) (room *entity.Room, err error) {
	err = c.txManager.WithTransaction(ctx, func(ctx context.Context) error {
		room := &entity.Room{
			Name:        request.Name,
			Description: request.Description,
		}
		return c.repo.Create(ctx, room)
	})

	return

}

func (roomController *roomControllerImpl) Update(ctx context.Context, request RoomRequest) error {
	return nil
}

func (roomController *roomControllerImpl) Delete(ctx context.Context, roomToken string) error {
	return nil
}

func ProvideRoomController(txManager database.TransactionManger, roomRepo entity.RoomRepository) RoomController {
	return &roomControllerImpl{
		txManager: txManager,
		repo:      roomRepo,
	}
}
