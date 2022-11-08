package entity

type Room struct {
	Token       int64  `db:"token"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

var Schema = `
	CREATE TABLE room (
		token 		bigint,
		name 		text,
		description text
	);
`
