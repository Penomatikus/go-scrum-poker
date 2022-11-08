package entity

type Participant struct {
	Name      string `db:"name"`
	RoomToken int64  `db:"room_token"`
}

var Schema = `
	CREATE TABLE participant (
		name 			text,
		room_token		bigint
	);
`
