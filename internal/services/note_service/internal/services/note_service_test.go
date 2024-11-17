package services

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/peygy/medapp/internal/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestInsertUserNote_Success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	log, _ := logger.NewLogger()

	ns := NewNoteService(db, log)

	mock.ExpectExec(`INSERT INTO appointments`).
		WithArgs("1", "Dr. Smith", "Cardiology", "2024-11-18").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := ns.InsertUserNote("1", "Dr. Smith", "Cardiology", "2024-11-18")

	assert.NoError(t, err)
}

func TestInsertUserNote_UserIdEmpty(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()
	log, _ := logger.NewLogger()

	ns := NewNoteService(db, log)

	err := ns.InsertUserNote("", "Dr. Smith", "Cardiology", "2024-11-18")

	assert.Error(t, err)
	assert.Equal(t, "note-services: user id is empty", err.Error())
}

func TestInsertUserNote_DBError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	log, _ := logger.NewLogger()

	ns := NewNoteService(db, log)

	mock.ExpectExec(`INSERT INTO appointments`).
		WithArgs("1", "Dr. Smith", "Cardiology", "2024-11-18").
		WillReturnError(errors.New("db error"))

	err := ns.InsertUserNote("1", "Dr. Smith", "Cardiology", "2024-11-18")

	assert.Error(t, err)
	assert.Equal(t, "note-services: can't inserts new note data to database", err.Error())
}

func TestGetUserNotes_Success(t *testing.T) {
	// Mock базы данных и логгера
	db, mock, _ := sqlmock.New()
	defer db.Close()
	log, _ := logger.NewLogger()

	ns := NewNoteService(db, log)

	// Ожидаем, что запрос вернет 1 запись
	rows := sqlmock.NewRows([]string{"id", "doctor_name", "specialization", "visit_date"}).
		AddRow("1", "Dr. Smith", "Cardiology", "2024-11-18")

	mock.ExpectQuery(`SELECT id, doctor_name, specialization, visit_date`).
		WithArgs("1").
		WillReturnRows(rows)

	// Вызов метода
	notes, err := ns.GetUserNotes("1")

	// Проверка
	assert.NoError(t, err)
	assert.Len(t, notes, 1)
	assert.Equal(t, "Dr. Smith", notes[0].Doctor_name)
}

func TestGetUserNotes_UserIdEmpty(t *testing.T) {
	// Mock базы данных и логгера
	db, _, _ := sqlmock.New()
	defer db.Close()
	log, _ := logger.NewLogger()

	ns := NewNoteService(db, log)

	// Вызов метода с пустым userId
	notes, err := ns.GetUserNotes("")

	// Проверка
	assert.Error(t, err)
	assert.Equal(t, "note-services: user id is empty", err.Error())
	assert.Len(t, notes, 0)
}

func TestGetUserNotes_DBError(t *testing.T) {
	// Mock базы данных и логгера
	db, mock, _ := sqlmock.New()
	defer db.Close()
	log, _ := logger.NewLogger()

	ns := NewNoteService(db, log)

	// Ожидаем, что запрос вернет ошибку
	mock.ExpectQuery(`SELECT id, doctor_name, specialization, visit_date`).
		WithArgs("1").
		WillReturnError(errors.New("db error"))

	// Вызов метода
	notes, err := ns.GetUserNotes("1")

	// Проверка
	assert.Error(t, err)
	assert.Equal(t, "note-services: can't get notes data from database", err.Error())
	assert.Len(t, notes, 0)
}

func TestGetUserNotes_ScanError(t *testing.T) {
	// Mock базы данных и логгера
	db, mock, _ := sqlmock.New()
	defer db.Close()
	log, _ := logger.NewLogger()

	ns := NewNoteService(db, log)

	// Ожидаем, что запрос вернет некорректные данные для сканирования
	rows := sqlmock.NewRows([]string{"id", "doctor_name", "specialization", "visit_date"}).
		AddRow("1", "Dr. Smith", "Cardiology", nil)

	mock.ExpectQuery(`SELECT id, doctor_name, specialization, visit_date`).
		WithArgs("1").
		WillReturnRows(rows)

	// Вызов метода
	notes, err := ns.GetUserNotes("1")

	// Проверка
	assert.Error(t, err)
	assert.Equal(t, "note-services: error scanning note data", err.Error())
	assert.Len(t, notes, 0)
}
