package main

import (
	"context"
	"github.com/joho/godotenv"
	"go-torent-parse/internal/config"
	"go-torent-parse/internal/storage/postgresql"
)

func main() {
	initEnv()

	cfg := config.MustLoad()

	storage, err := postgresql.New(context.Background(), cfg.DB)
	if err != nil {
		panic(err.Error())
	}
}

func initEnv() {
	if err := godotenv.Load(); err != nil {
		panic("env file does not loading: " + err.Error())
	}
}
