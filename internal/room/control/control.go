package control

import (
	"context"

	"github.com/penomatikus/go-scrum-poker/internal/room/entity"
)

type (
	RoomCreateRequest struct {
		Token       string
		Name        string
		Description string
	}

	RoomUpdateRequest struct {
		RoomCreateRequest
	}
)

type (
	RoomController interface {
		Create(context.Context, RoomCreateRequest) (*entity.Room, error)
		Update(context.Context, RoomUpdateRequest) error
		Delete(context.Context, string) error
	}

	roomControllerImpl struct{}
)

var _ RoomController = &roomControllerImpl{}

func (roomController *roomControllerImpl) Create(ctx context.Context, request RoomCreateRequest) (*entity.Room, error) {
	return nil, nil
}

func (roomController *roomControllerImpl) Update(ctx context.Context, request RoomUpdateRequest) error {
	return nil
}

func (roomController *roomControllerImpl) Delete(ctx context.Context, roomToken string) error {
	return nil
}
