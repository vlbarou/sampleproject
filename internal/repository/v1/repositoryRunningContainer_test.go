package v1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vlbarou/sampleproject/internal/model"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// go test -v repositoryRunningContainer_test.go repository.go
// go test --cover repositoryRunningContainer_test.go repository.go
func TestInsert(t *testing.T) {
	// Connect to your MariaDB database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"user", "password", "localhost", "3306", "testdb")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Disable logging for tests
	})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the User model
	if err := db.AutoMigrate(&model.User{}); err != nil {
		t.Fatalf("Failed to migrate User model: %v", err)
	}

	// Create a new user
	user := model.User{
		Name:     "John Doe",
		Username: "JohnDoe",
		Email:    "doe@example.com",
	}

	repo := NewUserRepository(db)
	err = repo.CreateUser(&user)

	// Assert that the user was inserted correctly
	if user.ID == 0 {
		t.Error("User ID should not be 0 after insertion")
	}

	// You can also check for specific values in the database
	var insertedUser model.User
	if err := db.First(&insertedUser, user.ID).Error; err != nil {
		t.Fatalf("Failed to find inserted user: %v", err)
	}

	if insertedUser.Name != user.Name {
		t.Errorf("Inserted user values do not match expected values")
	}

	t.Logf("User inserted successfully with ID: %d", user.ID)
}

func TestGetUserById(t *testing.T) {
	// Connect to your MariaDB database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"user", "password", "localhost", "3306", "testdb")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Disable logging for tests
	})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the User model
	if err := db.AutoMigrate(&model.User{}); err != nil {
		t.Fatalf("Failed to migrate User model: %v", err)
	}

	// Create a new user
	user := model.User{
		Name:     "John Doe",
		Username: "JohnDoe",
		Email:    "doe@example.com",
	}

	repo := NewUserRepository(db)
	err = repo.CreateUser(&user)

	// Assert that the user was inserted correctly
	if user.ID == 0 {
		t.Error("User ID should not be 0 after insertion")
	}

	// You can also check for specific values in the database
	retrievedUser, err := repo.GetUserByID(user.ID)
	if err != nil {
		t.Fatalf("Failed to find inserted user: %v", err)
	}

	assert.Equal(t, user.Name, retrievedUser.Name)
	assert.Equal(t, user.Username, retrievedUser.Username)
	assert.Equal(t, user.Email, retrievedUser.Email)

	t.Logf("User inserted successfully with ID: %d", user.ID)
}

func TestGetAllUsers1(t *testing.T) {
	// Connect to your MariaDB database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"user", "password", "localhost", "3306", "testdb")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Disable logging for tests
	})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the User model
	if err := db.AutoMigrate(&model.User{}); err != nil {
		t.Fatalf("Failed to migrate User model: %v", err)
	}

	// Create a new user
	user1 := model.User{
		Name:     "John Doe",
		Username: "JohnDoe",
		Email:    "doe@example.com",
	}

	repo := NewUserRepository(db)
	err = repo.CreateUser(&user1)

	user2 := model.User{
		Name:     "Vlasis Bar",
		Username: "VlasisBar",
		Email:    "vlasis@example.com",
	}
	err = repo.CreateUser(&user2)

	// Assert that the user was inserted correctly
	if user1.ID == 0 {
		t.Error("User ID should not be 0 after insertion")
	}

	if user2.ID == 0 {
		t.Error("User ID should not be 0 after insertion")
	}

	// You can also check for specific values in the database
	retrievedUsers, err := repo.GetAllUsers()
	if err != nil {
		t.Fatalf("Failed to find inserted user: %v", err)
	}

	assert.True(t, len(retrievedUsers) > 1) // Why not asserting for len(retrievedUsers) == 2 ?
}
