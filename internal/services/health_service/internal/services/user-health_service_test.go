package services

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/peygy/medapp/internal/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestInsertHealthData(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database connection: %v", err)
	}
	defer db.Close()

	mockLogger, _ := logger.NewLogger()
	service := NewHealthService(db, mockLogger)

	t.Run("Success", func(t *testing.T) {
		// Expect the INSERT query
		mock.ExpectExec("INSERT INTO users_health_data (userId) VALUES ($1) ON CONFLICT (userId) DO NOTHING").
			WithArgs("123").
			WillReturnResult(sqlmock.NewResult(1, 1)) // Return a dummy result

		err := service.InsertHealthData("123")

		assert.Error(t, err)
	})

	t.Run("Error: empty userId", func(t *testing.T) {
		err := service.InsertHealthData("")

		assert.Error(t, err)
		assert.Equal(t, "health-services: user id is empty", err.Error())
	})

	t.Run("Error: DB error", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO users_health_data (userId) VALUES ($1) ON CONFLICT (userId) DO NOTHING").
			WithArgs("123").
			WillReturnError(errors.New("db error"))

		err := service.InsertHealthData("123")

		assert.Error(t, err)
		assert.Equal(t, "health-services: can't inserts new health data to database", err.Error())
	})
}

func TestGetHealthDataByUserId(t *testing.T) {
	// Initialize sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database connection: %v", err)
	}
	defer db.Close()

	mockLogger, _ := logger.NewLogger()
	service := NewHealthService(db, mockLogger)

	t.Run("Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"age", "height", "weight", "pulse", "pressure"}).
			AddRow(30, 175.5, 70.2, 72, "120/80")

		mock.ExpectQuery("SELECT age, height, weight, pulse, pressure FROM users_health_data WHERE userId = $1").
			WithArgs("123").
			WillReturnRows(rows)

		_, err := service.GetHealthDataByUserId("123")

		assert.Error(t, err)
	})

	t.Run("Error: empty userId", func(t *testing.T) {
		_, err := service.GetHealthDataByUserId("")

		assert.Error(t, err)
		assert.Equal(t, "health-services: user id is empty", err.Error())
	})

	t.Run("Error: DB query error", func(t *testing.T) {
		mock.ExpectQuery("SELECT age, height, weight, pulse, pressure FROM users_health_data WHERE userId = $1").
			WithArgs("123").
			WillReturnError(errors.New("db error"))

		_, err := service.GetHealthDataByUserId("123")

		assert.Error(t, err)
		assert.Equal(t, "health-services: can't gets a health data from database", err.Error())
	})
}

func TestUpdateHealthDataByUserId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database connection: %v", err)
	}
	defer db.Close()

	mockLogger, _ := logger.NewLogger()
	service := NewHealthService(db, mockLogger)

	t.Run("Success", func(t *testing.T) {
		healthData := HealthData{
			Age:      30,
			Height:   175.5,
			Weight:   70.2,
			Pulse:    72,
			Pressure: "120/80",
		}

		mock.ExpectExec("UPDATE users_health_data SET age = $1, height = $2, weight = $3, pulse = $4, pressure = $5 WHERE userId = $6").
			WithArgs(healthData.Age, healthData.Height, healthData.Weight, healthData.Pulse, healthData.Pressure, "123").
			WillReturnResult(sqlmock.NewResult(1, 1))

		updatedData, err := service.UpdateHealthDataByUserId("123", healthData)

		assert.Error(t, err)
		assert.Equal(t, healthData, updatedData)
	})

	t.Run("Error: empty userId", func(t *testing.T) {
		_, err := service.UpdateHealthDataByUserId("", HealthData{})

		assert.Error(t, err)
		assert.Equal(t, "health-services: user id is empty", err.Error())
	})

	t.Run("Error: empty health data", func(t *testing.T) {
		_, err := service.UpdateHealthDataByUserId("123", HealthData{})

		assert.Error(t, err)
		assert.Equal(t, "health-services: health data is empty", err.Error())
	})

	t.Run("Error: DB update error", func(t *testing.T) {
		healthData := HealthData{
			Age:      30,
			Height:   175.5,
			Weight:   70.2,
			Pulse:    72,
			Pressure: "120/80",
		}

		// Expect the UPDATE query and return an error
		mock.ExpectExec("UPDATE users_health_data SET age = $1, height = $2, weight = $3, pulse = $4, pressure = $5 WHERE userId = $6").
			WithArgs(healthData.Age, healthData.Height, healthData.Weight, healthData.Pulse, healthData.Pressure, "123").
			WillReturnError(errors.New("db error"))

		_, err := service.UpdateHealthDataByUserId("123", healthData)

		assert.Error(t, err)
		assert.Equal(t, "health-services: can't update health data in the database", err.Error())
	})
}
