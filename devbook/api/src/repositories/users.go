package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (u users) Create(usuario models.User) (uint64, error) {
	statement, erro := u.db.Prepare(
		"INSERT INTO users (name, nick, email, passkeyword) VALUES (?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(usuario.Name, usuario.Nick, usuario.Email, usuario.Password)
	if erro != nil {
		return 0, erro
	}

	lastID, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastID), nil
}
