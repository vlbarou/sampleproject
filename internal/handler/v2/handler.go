package v2

import (
	"encoding/json"
	"github.com/vlbarou/sampleproject/internal/repository/v2"
	"github.com/vlbarou/sampleproject/pkg/validator"
	"net/http"

	"github.com/vlbarou/sampleproject/internal/model"
)

type UserHandler struct {
	repo v2.UserRepository
}

type HealthResponse struct {
	Health string `json:"message"`
}

func NewUserHandler(repo v2.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// logic to retrieve the health status, typically implemented at service layer

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthResponse{Health: "running"})
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.repo.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	user, err := h.repo.GetUserByID(2)
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if !validator.ValidateUsername(user.Username) {
		http.Error(w, "Invalid username", http.StatusBadRequest)
		return
	}

	if !validator.ValidateEmail(user.Email) {
		http.Error(w, "Invalid email", http.StatusBadRequest)
		return
	}

	if err := h.repo.CreateUser(&user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("post was successful")
}
