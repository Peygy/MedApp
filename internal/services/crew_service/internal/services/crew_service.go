package services

import (
	"database/sql"
	"errors"

	"github.com/peygy/medapp/internal/pkg/logger"
)

type CrewData struct {
	DoctorId        string
	DoctorName      string
	Specialization  string
	ExperienceYears int
}

type ICrewService interface {
	GetAllCrew() ([]CrewData, error)
}

type crewService struct {
	db  *sql.DB
	log logger.ILogger
}

func NewCrewService(db *sql.DB, log logger.ILogger) ICrewService {
	log.Info("CrewService created")
	return &crewService{
		db:  db,
		log: log,
	}
}

func (cs *crewService) GetAllCrew() ([]CrewData, error) {
	var crews []CrewData

	query := `SELECT id, doctor_name, specialization, experience_years FROM doctors`
	rows, err := cs.db.Query(query)
	if err != nil {
		cs.log.Errorf("Can't get crews data from database: %v", err)
		return crews, errors.New("crew-services: can't get crews data from database")
	}
	defer rows.Close()

	for rows.Next() {
		var crewData CrewData
		err := rows.Scan(
			&crewData.DoctorId,
			&crewData.DoctorName,
			&crewData.Specialization,
			&crewData.ExperienceYears,
		)
		if err != nil {
			cs.log.Errorf("Error scanning crew data: %v", err)
			return crews, errors.New("crew-services: error scanning crews data")
		}
		crews = append(crews, crewData)
	}

	if err = rows.Err(); err != nil {
		cs.log.Errorf("Error occurred during rows iteration: %v", err)
		return crews, errors.New("crew-services: error during rows iteration")
	}

	return crews, nil
}
