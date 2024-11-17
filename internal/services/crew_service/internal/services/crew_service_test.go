package services

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/peygy/medapp/internal/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestGetAllCrew_Success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	log, _ := logger.NewLogger()

	cs := NewCrewService(db, log)

	rows := sqlmock.NewRows([]string{"id", "doctor_name", "specialization", "experience_years"}).
		AddRow("1", "Dr. Smith", "Cardiology", 10)

	mock.ExpectQuery(`SELECT id, doctor_name, specialization, experience_years`).
		WillReturnRows(rows)

	crews, err := cs.GetAllCrew()

	assert.NoError(t, err)
	assert.Len(t, crews, 1)
	assert.Equal(t, "Dr. Smith", crews[0].DoctorName)
	assert.Equal(t, 10, crews[0].ExperienceYears)
}

func TestGetAllCrew_DBError(t *testing.T) {
	// Mock базы данных и логгера
	db, mock, _ := sqlmock.New()
	defer db.Close()
	log, _ := logger.NewLogger()

	cs := NewCrewService(db, log)

	// Ожидаем, что запрос вернет ошибку
	mock.ExpectQuery(`SELECT id, doctor_name, specialization, experience_years`).
		WillReturnError(errors.New("db error"))

	// Вызов метода
	crews, err := cs.GetAllCrew()

	// Проверка
	assert.Error(t, err)
	assert.Equal(t, "crew-services: can't get crews data from database", err.Error())
	assert.Len(t, crews, 0)
}

func TestGetAllCrew_ScanError(t *testing.T) {
	// Mock базы данных и логгера
	db, mock, _ := sqlmock.New()
	defer db.Close()
	log, _ := logger.NewLogger()

	cs := NewCrewService(db, log)

	// Ожидаем, что запрос вернет некорректные данные для сканирования
	rows := sqlmock.NewRows([]string{"id", "doctor_name", "specialization", "experience_years"}).
		AddRow("1", "Dr. Smith", "Cardiology", nil) // ExperienceYears = nil

	mock.ExpectQuery(`SELECT id, doctor_name, specialization, experience_years`).
		WillReturnRows(rows)

	// Вызов метода
	crews, err := cs.GetAllCrew()

	// Проверка
	assert.Error(t, err)
	assert.Equal(t, "crew-services: error scanning crews data", err.Error())
	assert.Len(t, crews, 0)
}

func TestGetAllCrew_EmptyResult(t *testing.T) {
	// Mock базы данных и логгера
	db, mock, _ := sqlmock.New()
	defer db.Close()
	log, _ := logger.NewLogger()

	cs := NewCrewService(db, log)

	// Ожидаем, что запрос вернет пустые строки
	rows := sqlmock.NewRows([]string{"id", "doctor_name", "specialization", "experience_years"})

	mock.ExpectQuery(`SELECT id, doctor_name, specialization, experience_years`).
		WillReturnRows(rows)

	// Вызов метода
	crews, err := cs.GetAllCrew()

	// Проверка
	assert.NoError(t, err)
	assert.Len(t, crews, 0)
}
