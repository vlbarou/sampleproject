package v2

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vlbarou/sampleproject/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"regexp"
	"testing"
)

func setupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true, // Skips version check for MariaDB
	}), &gorm.Config{
		SkipDefaultTransaction: true,                                // Disables implicit transactions. In this case there is no need to add mock.ExpectBegin() and mock.ExpectCommit() around insert commands
		Logger:                 logger.Default.LogMode(logger.Info), // Enable SQL logging
	})

	return gormDB, mock, err
}

func TestInsertUser(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)
	defer db.DB()

	repo := NewUserRepository(db)

	// Mock transaction BEGIN
	//mock.ExpectBegin()
	// Mock expected SQL insert query
	// INSERT INTO `users` (`name`,`username`,`email`) VALUES ('John Doe','JohnDoe','doe@example.com') RETURNING `id`
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`name`,`username`,`email`) VALUES (?,?,?)")).
		WithArgs("John Doe", "JohnDoe", "doe@example.com").
		WillReturnResult(sqlmock.NewResult(1, 1)) // Simulating last insert ID = 1

	// Mock transaction COMMIT
	//mock.ExpectCommit()

	// Create a new user
	users := &model.User{
		Name:     "John Doe",
		Username: "JohnDoe",
		Email:    "doe@example.com",
	}

	// Call the InsertUser function
	err = repo.CreateUser(users)

	// Assertions
	assert.NoError(t, err)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUserByID(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)
	defer db.DB()

	repo := NewUserRepository(db)

	// Mock data
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` = ? ORDER BY `users`.`id` LIMIT ?")).
		WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "username", "email"}).AddRow(1, "John Doe", "JohnDoe", "doe@example.com"))

	// Call function
	user, err := repo.GetUserByID(1)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "John Doe", user.Name)

	// Ensure expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllUsers2(t *testing.T) {
	db, mock, err := setupMockDB()
	assert.NoError(t, err)
	defer db.DB()

	repo := NewUserRepository(db)

	// Mock data
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "username", "email"}).
			AddRow(1, "John Doe", "JohnDoe", "doe@example.com").
			AddRow(2, "John Doe", "JohnDoe", "doe@example.com"))

	// Call function
	users, err := repo.GetAllUsers()

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, users)

	// Ensure expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Equal(t, 2, len(users))
}
