package v2

// https://chatgpt.com/share/67d612e5-21b4-8009-9c5f-f02fec6e375b
import (
	"github.com/stretchr/testify/assert"
	"github.com/vlbarou/sampleproject/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	mockrepository := new(MockUserRepository)
	handler := &UserHandler{repo: mockrepository}

	users := []model.User{
		{
			Name:     "John Doe",
			Username: "JohnDoe",
			Email:    "doe@example.com",
		},
		{
			Name:     "Vlasis bar",
			Username: "VlasisBar",
			Email:    "vlasis@example.com",
		},
	}

	mockrepository.On("GetAllUsers").Return(users, nil)

	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetUsers(rr, req)

	// Assertions
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `[{"ID":0,"Name":"John Doe","Username":"JohnDoe","Email":"doe@example.com"},{"ID":0,"Name":"Vlasis bar","Username":"VlasisBar","Email":"vlasis@example.com"}]`, rr.Body.String())

	mockrepository.AssertExpectations(t)
}

//func TestCreateUser(t *testing.T) {
//	mockRepo := &MockUserRepository{
//		MockCreateUser: func(user *model.User) error {
//			return nil
//		},
//	}
//
//	handler := &UserHandler{repo: mockRepo}
//
//	user := &model.User{
//		Name:     "John Doe",
//		Username: "JohnDoe",
//		Email:    "doe@example.com",
//	}
//
//	body, _ := json.Marshal(user)
//	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
//	assert.NoError(t, err)
//
//	rr := httptest.NewRecorder()
//	handler.CreateUser(rr, req)
//
//	assert.Equal(t, http.StatusCreated, rr.Code)
//	assert.JSONEq(t, "post was successful", rr.Body.String())
//}
