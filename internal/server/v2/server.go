package v2

import (
	"fmt"
	handlerV2 "github.com/vlbarou/sampleproject/internal/handler/v2"
	repoV2 "github.com/vlbarou/sampleproject/internal/repository/v2"
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
	userHandler := handlerV2.NewUserHandler(repoV2.NewUserRepository(s.db))

	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/user", userHandler.GetUserById).Methods("GET")
	r.HandleFunc("/user", userHandler.CreateUser).Methods("POST")

	addr := ":8080"
	fmt.Printf("Server running on %s\n", addr)
	return http.ListenAndServe(addr, r)
}
