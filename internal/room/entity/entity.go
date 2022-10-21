package entity

type Room struct {
	Token       string `db:"token"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

var Schema = `
	CREATE TABLE room (
		token 		text,
		name 		text,
		description text
	);
`
