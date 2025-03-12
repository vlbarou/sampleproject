package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"golang/config"
	"golang/internal/handler"
	"golang/internal/repository"
	"gorm.io/gorm"
)

type Server struct {
	config *config.Config
	db     *gorm.DB
}

func NewServer(cfg *config.Config, db *gorm.DB) *Server {
	return &Server{config: cfg, db: db}
}

func (s *Server) Start() error {
	r := mux.NewRouter()

	userRepo := repository.NewUserRepository(s.db)
	userHandler := handler.NewUserHandler(userRepo)

	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")

	addr := ":8080"
	fmt.Printf("Server running on %s\n", addr)
	return http.ListenAndServe(addr, r)
}
