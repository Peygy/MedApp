package managers

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/peygy/medapp/internal/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestNewRoleManager(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock db: %v", err)
	}
	defer db.Close()

	log, _ := logger.NewLogger()

	for _, role := range roles {
		mock.ExpectExec(`INSERT INTO roles (role_name) VALUES \($1\) ON CONFLICT \(role_name\) DO NOTHING`).
			WithArgs(role)
	}

	rm := NewRoleManager(db, log)
	assert.Nil(t, rm)
}

func TestAddRolesToUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock db: %v", err)
	}
	defer db.Close()

	log, _ := logger.NewLogger()

	rm := &roleManger{db: db, log: log}

	mock.ExpectQuery(`SELECT id FROM roles WHERE role_name = \$(\d+)`).
		WithArgs("admin").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	mock.ExpectExec(`INSERT INTO users_roles \(user_id, role_id\) VALUES \($1, $2\) ON CONFLICT DO NOTHING`).
		WithArgs("user1", "1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = rm.AddRolesToUser("user1", "admin")
	assert.Error(t, err)

	err = mock.ExpectationsWereMet()
	assert.Error(t, err)
}

func TestGetUserRole(t *testing.T) {
	// Create a mock DB and logger
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock db: %v", err)
	}
	defer db.Close()

	log, _ := logger.NewLogger() // Assuming a mock logger, or use a real one

	rm := &roleManger{db: db, log: log}

	// Mock the user-role query
	mock.ExpectQuery(`SELECT role_id FROM users_roles WHERE user_id = \$(\d+)`).
		WithArgs("user1").
		WillReturnRows(sqlmock.NewRows([]string{"role_id"}).AddRow("1"))

	// Mock the role query
	mock.ExpectQuery(`SELECT role_name FROM roles WHERE id = \$(\d+)`).
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"role_name"}).AddRow("admin"))

	// Test getting user role
	role, err := rm.GetUserRole("user1")
	assert.NoError(t, err)
	assert.Equal(t, "admin", role)

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestDeleteRolesFromUser(t *testing.T) {
	// Create a mock DB and logger
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock db: %v", err)
	}
	defer db.Close()

	log, _ := logger.NewLogger() // Assuming a mock logger, or use a real one

	rm := &roleManger{db: db, log: log}

	// Set up mock for checking role existence
	mock.ExpectQuery(`SELECT id FROM roles WHERE role_name = \$(\d+)`).
		WithArgs("admin").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	// Mock the delete from users_roles
	mock.ExpectExec(`DELETE FROM users_roles WHERE user_id = \$(\d+) AND role_id = \$(\d+)`).
		WithArgs("user1", "1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Test deleting roles from a user
	err = rm.DeleteRolesFromUser("user1", "admin")
	assert.Error(t, err)

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.Error(t, err)
}

func TestRoleNotFound(t *testing.T) {
	// Create a mock DB and logger
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock db: %v", err)
	}
	defer db.Close()

	log, _ := logger.NewLogger() // Assuming a mock logger, or use a real one

	rm := &roleManger{db: db, log: log}

	// Mock the role existence check for a non-existent role
	mock.ExpectQuery(`SELECT id FROM roles WHERE role_name = \$(\d+)`).
		WithArgs("nonexistent").
		WillReturnError(sql.ErrNoRows)

	// Test adding a nonexistent role
	err = rm.AddRolesToUser("user1", "nonexistent")
	assert.Error(t, err)

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	assert.Error(t, err)
}
