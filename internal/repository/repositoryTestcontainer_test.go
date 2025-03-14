package repository

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/vlbarou/sampleproject/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"testing"
)

var testDB *gorm.DB

func setupTestDB() (*gorm.DB, func()) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "mariadb:latest",
		ExposedPorts: []string{"3306/tcp"},
		Env: map[string]string{
			"MARIADB_ROOT_PASSWORD": "root",
			"MARIADB_DATABASE":      "testdb",
			"MARIADB_USER":          "user",
			"MARIADB_PASSWORD":      "password",
		},
		WaitingFor: wait.ForLog("port: 3306  MariaDB Server"),
	}

	mariaDBContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		fmt.Printf("Failed to start MariaDB container: %v", err)
	}

	host, err := mariaDBContainer.Host(ctx)
	if err != nil {
		fmt.Printf("Failed to get container host: %v", err)
	}

	port, err := mariaDBContainer.MappedPort(ctx, "3306")
	if err != nil {
		fmt.Printf("Failed to get container port: %v", err)
	}

	dsn := fmt.Sprintf("user:password@tcp(%s:%s)/testdb?charset=utf8mb4&parseTime=True&loc=Local", host, port.Port())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable SQL loggin
	})
	if err != nil {
		log.Fatalf("Failed to connect to MariaDB: %v", err)
	}

	// Cleanup function to stop the container
	teardown := func() {
		mariaDBContainer.Terminate(ctx)
	}

	return db, teardown
}

func TestCreateUser(t *testing.T) {
	db, cleanup := setupTestDB()
	defer cleanup()

	err := db.AutoMigrate(model.User{})
	if err != nil {
		fmt.Println("automigration failed")
	} else {
		fmt.Println("automigration OK")
	}

	user := model.User{
		Name:     "John Doe",
		Username: "JohnDoe",
		Email:    "doe@example.com",
	}

	repo := NewUserRepository(db)
	err = repo.CreateUser(&user)
	if err != nil {
		fmt.Println("failed to create user")
	} else {
		fmt.Println("user created")
	}

	var result model.User
	db.First(&result, "name = ?", "John Doe")

	assert.Equal(t, "John Doe", result.Name)
	assert.Equal(t, "JohnDoe", result.Username)
	assert.Equal(t, "doe@example.com", result.Email)
}

func TestGetUserById1(t *testing.T) {
	db, cleanup := setupTestDB()
	defer cleanup()

	err := db.AutoMigrate(model.User{})
	if err != nil {
		fmt.Println("automigration failed")
	} else {
		fmt.Println("automigration OK")
	}

	user := model.User{
		Name:     "John Doe",
		Username: "JohnDoe",
		Email:    "doe@example.com",
	}

	repo := NewUserRepository(db)
	err = repo.CreateUser(&user)
	if err != nil {
		fmt.Println("failed to create user")
	} else {
		fmt.Println("user created")
	}

	retrievedUser, err := repo.GetUserByID(user.ID)
	if err != nil {
		t.Fatalf("Failed to find inserted user: %v", err)
	}

	assert.Equal(t, user.Name, retrievedUser.Name)
	assert.Equal(t, user.Username, retrievedUser.Username)
	assert.Equal(t, user.Email, retrievedUser.Email)
}

func TestGetAaaUsers1(t *testing.T) {
	db, cleanup := setupTestDB()
	defer cleanup()

	err := db.AutoMigrate(model.User{})
	if err != nil {
		fmt.Println("automigration failed")
	} else {
		fmt.Println("automigration OK")
	}

	repo := NewUserRepository(db)

	user1 := model.User{
		Name:     "John Doe",
		Username: "JohnDoe",
		Email:    "doe@example.com",
	}

	user2 := model.User{
		Name:     "Vlasis Bar",
		Username: "VlasisBar",
		Email:    "vlasis@example.com",
	}

	err = repo.CreateUser(&user1)
	if err != nil {
		fmt.Println("failed to create user")
	} else {
		fmt.Println("user created")
	}

	err = repo.CreateUser(&user2)
	if err != nil {
		fmt.Println("failed to create user")
	} else {
		fmt.Println("user created")
	}

	retrievedUsers, err := repo.GetAllUsers()
	if err != nil {
		t.Fatalf("Failed to find inserted user: %v", err)
	}

	assert.Equal(t, 2, len(retrievedUsers))
}
