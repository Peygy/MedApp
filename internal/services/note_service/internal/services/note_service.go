package services

import (
	"database/sql"
	"errors"

	"github.com/peygy/medapp/internal/pkg/logger"
)

type INoteService interface {
	InsertUserNote(userId, doctorName, specialization, visitDate string) error
	GetUserNotes(userId string) ([]NoteData, error)
}

type noteService struct {
	db  *sql.DB
	log logger.ILogger
}

type NoteData struct {
	Id             string
	Doctor_name    string
	Specialization string
	Visit_date     string
}

func NewNoteService(db *sql.DB, log logger.ILogger) INoteService {
	log.Info("NoteService created")
	return &noteService{
		db:  db,
		log: log,
	}
}

func (ns *noteService) InsertUserNote(userId, doctorName, specialization, visitDate string) error {
	if userId == "" {
		ns.log.Errorf("User id (%s) is empty", userId)
		return errors.New("note-services: user id is empty")
	}

	query := `INSERT INTO appointments (userId, doctor_name, specialization, visit_date) VALUES ($1, $2, $3, $4)`
	if _, err := ns.db.Exec(query, userId, doctorName, specialization, visitDate); err != nil {
		ns.log.Errorf("Can't inserts new note data for user (%s) to database: %v", userId, err)
		return errors.New("note-services: can't inserts new note data to database")
	}

	ns.log.Info("New health data inserted successfully")
	return nil
}

func (ns *noteService) GetUserNotes(userId string) ([]NoteData, error) {
	var notes []NoteData

	if userId == "" {
		ns.log.Errorf("User id (%s) is empty", userId)
		return notes, errors.New("note-services: user id is empty")
	}

	query := `SELECT id, doctor_name, specialization, visit_date FROM appointments WHERE userId = $1`
	rows, err := ns.db.Query(query, userId)
	if err != nil {
		ns.log.Errorf("Can't get notes data for user (%s) from database: %v", userId, err)
		return notes, errors.New("note-services: can't get notes data from database")
	}
	defer rows.Close()

	for rows.Next() {
		var noteData NoteData
		err := rows.Scan(
			&noteData.Id,
			&noteData.Doctor_name,
			&noteData.Specialization,
			&noteData.Visit_date,
		)
		if err != nil {
			ns.log.Errorf("Error scanning note data for user (%s): %v", userId, err)
			return notes, errors.New("note-services: error scanning note data")
		}
		notes = append(notes, noteData)
	}

	if err = rows.Err(); err != nil {
		ns.log.Errorf("Error occurred during rows iteration for user (%s): %v", userId, err)
		return notes, errors.New("note-services: error during rows iteration")
	}

	ns.log.Infof("User (%s) notes data retrieved successfully", userId)
	return notes, nil
}
