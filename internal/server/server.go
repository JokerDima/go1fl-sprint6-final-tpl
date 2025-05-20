package server

import (
	"log"
	"net/http"
	"time"

	"go1fl-sprint6-final-tpl/internal/handlers"

	"github.com/go-chi/chi/v5"
)

// Структура сервера
type Server struct {
	logger *log.Logger
	server *http.Server
}

// Запуск сервера
func NewServer(logger *log.Logger) *Server {
	router := chi.NewRouter()
	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	srvr := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger: logger,
		server: srvr,
	}

}

// Метод для запуска сервера
func (s *Server) Start() error {
	return s.server.ListenAndServe()
}
