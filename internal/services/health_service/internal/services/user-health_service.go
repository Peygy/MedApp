package services

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/peygy/medapp/internal/pkg/logger"
)

type HealthData struct {
	Age      int32
	Height   float32
	Weight   float32
	Pulse    int32
	Pressure string
}

type IHealthService interface {
	InsertHealthData(userId string) error
	GetHealthDataByUserId(userId string) (HealthData, error)
	UpdateHealthDataByUserId(userId string, healthData HealthData) (HealthData, error)
}

type healthService struct {
	db  *sql.DB
	log logger.ILogger
}

func NewHealthService(db *sql.DB, log logger.ILogger) IHealthService {
	log.Info("HealthService created")
	return &healthService{
		db:  db,
		log: log,
	}
}

func (hs *healthService) InsertHealthData(userId string) error {
	if userId == "" {
		hs.log.Errorf("User id (%s) is empty", userId)
		return errors.New("health-services: user id is empty")
	}

	query := `INSERT INTO users_health_data (userId) VALUES ($1) ON CONFLICT (userId) DO NOTHING`
	if _, err := hs.db.Exec(query, userId); err != nil {
		hs.log.Errorf("Can't inserts new health data for user (%s) to database: %v", userId, err)
		return errors.New("health-services: can't inserts new health data to database")
	}

	hs.log.Info("New health data inserted successfully")
	return nil
}

func (hs *healthService) GetHealthDataByUserId(userId string) (HealthData, error) {
	var healthData HealthData

	if userId == "" {
		hs.log.Errorf("User id (%s) is empty", userId)
		return healthData, errors.New("health-services: user id is empty")
	}

	query := `SELECT age, height, weight, pulse, pressure FROM users_health_data WHERE userId = $1`
	err := hs.db.QueryRow(query, userId).Scan(
		&healthData.Age,
		&healthData.Height,
		&healthData.Weight,
		&healthData.Pulse,
		&healthData.Pressure,
	)

	if err != nil {
		hs.log.Errorf("Can't gets a health data for user (%s) from database: %v", userId, err)
		return healthData, errors.New("health-services: can't gets a health data from database")
	}

	fmt.Printf("%+v\n", healthData)

	hs.log.Infof("User (%s) health data getted successfully", userId)
	return healthData, nil
}

func (hs *healthService) UpdateHealthDataByUserId(userId string, healthData HealthData) (HealthData, error) {
	if userId == "" {
		hs.log.Errorf("User id (%s) is empty", userId)
		return healthData, errors.New("health-services: user id is empty")
	}

	if (HealthData{}) == healthData {
		hs.log.Errorf("Health data is empty for user id: %s", userId)
		return healthData, errors.New("health-services: health data is empty")
	}

	query := `
		UPDATE users_health_data 
		SET age = $1, height = $2, weight = $3, pulse = $4, pressure = $5 
		WHERE userId = $6
		RETURNING age, height, weight, pulse, pressure`

	err := hs.db.QueryRow(query,
		healthData.Age,
		healthData.Height,
		healthData.Weight,
		healthData.Pulse,
		healthData.Pressure,
		userId).Scan(
		&healthData.Age,
		&healthData.Height,
		&healthData.Weight,
		&healthData.Pulse,
		&healthData.Pressure,
	)

	if err != nil {
		hs.log.Errorf("Can't update health data for user (%s) in the database: %v", userId, err)
		return healthData, errors.New("health-services: can't update health data in the database")
	}

	fmt.Printf("%+v\n", healthData)

	hs.log.Infof("User (%s) health data updated successfully", userId)
	return healthData, nil
}
