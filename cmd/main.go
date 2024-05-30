package main

import (
	"github.com/joho/godotenv"
	"go-torent-parse/internal/config"
)

func main() {
	initEnv()

	cfg := config.MustLoad()
}

func initEnv() {
	if err := godotenv.Load(); err != nil {
		panic("env file does not loading")
	}
}
