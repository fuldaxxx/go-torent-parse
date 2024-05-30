package postgresql

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go-torent-parse/internal/config"
	"time"
)

type Storage struct {
	db *sqlx.DB
}

func New(ctx context.Context, dbConfig config.DBConfig) (*Storage, error) {
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name)

	db, err := sqlx.ConnectContext(dbCtx, "postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to database: " + err.Error())
	}

	return &Storage{db: db}, nil
}
