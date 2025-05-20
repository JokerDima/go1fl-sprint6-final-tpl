package main

import (
	"log"
	"os"

	"go1fl-sprint6-final-tpl/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	r := server.NewServer(logger)
	if err := r.Start(); err != nil {
		logger.Fatalf("Ошибка при запуске сервера: %s", err)
		return
	}
}
