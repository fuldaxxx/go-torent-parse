package postgresql

import "github.com/jmoiron/sqlx"

type Storage struct {
	db *sqlx.DB
}
