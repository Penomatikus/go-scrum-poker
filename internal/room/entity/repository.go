package entity

type RoomToken string

type RoomCreateRequest struct {
	Token       string
	Name        string
	Description string
}

type RoomUpdateRequest struct {
	RoomCreateRequest
}

type RoomRepository interface {
	Create(RoomCreateRequest) (*Room, error)
	Update(RoomUpdateRequest) error
	Delete(RoomToken) error
}
