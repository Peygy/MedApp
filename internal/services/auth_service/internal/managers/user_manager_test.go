package managers

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/peygy/medapp/internal/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestInsertUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	log, _ := logger.NewLogger()
	manager := NewUserManager(db, log)

	user := UserRecord{
		UserName: "test_user",
		Password: "hashed_password",
	}

	mock.ExpectQuery("INSERT INTO users").
		WithArgs(user.UserName, user.Password).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	userId, err := manager.InsertUser(user)
	assert.NoError(t, err)
	assert.Equal(t, "1", userId)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestInsertUser_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	log, _ := logger.NewLogger()
	manager := NewUserManager(db, log)

	user := UserRecord{
		UserName: "test_user",
		Password: "hashed_password",
	}

	mock.ExpectQuery("INSERT INTO users").
		WithArgs(user.UserName, user.Password).
		WillReturnError(errors.New("insert error"))

	userId, err := manager.InsertUser(user)
	assert.Error(t, err)
	assert.Empty(t, userId)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetUserByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	log, _ := logger.NewLogger()
	manager := NewUserManager(db, log)

	expectedUser := UserRecord{
		Id:       "1",
		UserName: "test_user",
		Password: "hashed_password",
	}

	mock.ExpectQuery("SELECT id, username, password_hash FROM users WHERE username = \\$1").
		WithArgs("test_user").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password_hash"}).
			AddRow(expectedUser.Id, expectedUser.UserName, expectedUser.Password))

	user, err := manager.GetUserByName("test_user")
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetUserById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	log, _ := logger.NewLogger()
	manager := NewUserManager(db, log)

	expectedUser := UserRecord{
		Id:       "1",
		UserName: "test_user",
		Password: "hashed_password",
	}

	mock.ExpectQuery("SELECT id, username, password_hash FROM users WHERE id = \\$1").
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password_hash"}).
			AddRow(expectedUser.Id, expectedUser.UserName, expectedUser.Password))

	user, err := manager.GetUserById("1")
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetUserByName_EmptyUsername(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	log, _ := logger.NewLogger()
	manager := NewUserManager(db, log)

	user, err := manager.GetUserByName("")
	assert.Error(t, err)
	assert.Equal(t, UserRecord{}, user)
}
