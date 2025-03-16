package v2

// https://chatgpt.com/share/67d612e5-21b4-8009-9c5f-f02fec6e375b
import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/vlbarou/sampleproject/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllUsers(t *testing.T) {

	// arrange
	mockrepository := new(MockUserRepository)
	handler := &UserHandler{repo: mockrepository}

	users := []model.User{
		{
			ID:       1,
			Name:     "name1",
			Username: "username1",
			Email:    "email1@example.com",
		},
		{
			ID:       2,
			Name:     "name2",
			Username: "username2",
			Email:    "email2@example.com",
		},
	}

	mockrepository.On("GetAllUsers").Return(users, nil)

	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	// act
	handler.GetUsers(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `[{"ID":1,"Name":"name1","Username":"username1","Email":"email1@example.com"},{"ID":2,"Name":"name2","Username":"username2","Email":"email2@example.com"}]`, rr.Body.String())

	mockrepository.AssertExpectations(t)
}

func TestGetUserByID(t *testing.T) {

	// arrange
	mockrepository := new(MockUserRepository)
	handler := &UserHandler{repo: mockrepository}

	// Mock data
	user := model.User{
		ID:       1,
		Name:     "name1",
		Username: "username1",
		Email:    "email1@example.com",
	}

	mockrepository.On("GetUserByID", 1).Return(user, nil)

	req, err := http.NewRequest("GET", "/user?id=1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	// act
	handler.GetUserById(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"Email":"email1@example.com", "ID":1, "Name":"name1", "Username":"username1"}`, rr.Body.String())

	mockrepository.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {

	// arrange
	mockrepository := new(MockUserRepository)
	handler := &UserHandler{repo: mockrepository}

	user := model.User{
		ID:       1,
		Name:     "name1",
		Username: "username1",
		Email:    "email1@example.com",
	}

	body, _ := json.Marshal(user)

	// Expect CreateUser to be called with the given user and return nil (success)
	mockrepository.On("CreateUser", &user).Return(nil)

	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// act
	handler.CreateUser(rr, req)

	// Assert
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.JSONEq(t, `{"Message": "User created"}`, rr.Body.String())
	mockrepository.AssertExpectations(t) // Ensure all expectations are met
}
