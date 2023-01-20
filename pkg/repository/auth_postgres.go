package repository

import (
	app "github.com/HgCl2/rock_scissors_paper"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user app.Player) (int, error) {
	var id int
	query := `INSERT INTO users (fullname, username, password_hash) VALUES ($1, $2, $3)
		RETURNING id`

	row := r.db.QueryRow(query, user.Fullname, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
