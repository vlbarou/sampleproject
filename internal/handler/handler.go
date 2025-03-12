package handler

import (
	"encoding/json"
	"github.com/vlbarou/sampleproject/pkg/validator"
	"net/http"

	"github.com/vlbarou/sampleproject/internal/model"
	"github.com/vlbarou/sampleproject/internal/repository"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.repo.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	for _, user := range users {
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
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("post was successful")
}
