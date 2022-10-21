package entity

import (
	"context"
)

type RoomToken string

type RoomRepository interface {
	Create(context.Context, *Room) error
	Update(context.Context, Room) error
	Delete(context.Context, RoomToken) error
}
