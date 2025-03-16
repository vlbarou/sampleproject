package v1

import (
	"fmt"
	handlerV1 "github.com/vlbarou/sampleproject/internal/handler/v1"
	repoV1 "github.com/vlbarou/sampleproject/internal/repository/v1"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/vlbarou/sampleproject/config"
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

	userRepo := repoV1.NewUserRepository(s.db)
	userHandler := handlerV1.NewUserHandler(userRepo)

	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")

	addr := ":8080"
	fmt.Printf("Server running on %s\n", addr)
	return http.ListenAndServe(addr, r)
}
