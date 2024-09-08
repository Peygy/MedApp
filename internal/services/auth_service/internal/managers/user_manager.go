package managers

import (
	"database/sql"
	"errors"

	"github.com/peygy/medapp/internal/pkg/logger"
)

// Provides the service for managing user persistence store.
type IUserManager interface {
	// Adds new user to the database.
	// Returns new user id or error.
	InsertUser(user UserRecord) (string, error)
	// Gets user from the database by name.
	// Returns user model or error.
	GetUserByName(userName string) (UserRecord, error)
	// Gets user from the database by name.
	// Returns user model or error.
	GetUserById(userId string) (UserRecord, error)
}

type UserRecord struct {
	Id       string
	UserName string
	Password string
}

type userManger struct {
	db  *sql.DB
	log logger.ILogger
}

func NewUserManager(db *sql.DB, log logger.ILogger) IUserManager {
	log.Info("UserManager created")
	return &userManger{
		db:  db,
		log: log,
	}
}

func (um *userManger) InsertUser(user UserRecord) (string, error) {
	var userId string
	query := `INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING id`
	if err := um.db.QueryRow(query, user.UserName, user.Password).Scan(&userId); err != nil {
		um.log.Errorf("Can't adds new user (%s, %s) to the database: %v", user.UserName, user.Password, err)
		return "", errors.New("managers-user: can't creates user in the database")
	}

	um.log.Infof("User %s created successfully", userId)
	return userId, nil
}

func (um *userManger) GetUserByName(userName string) (UserRecord, error) {
	if userName == "" {
		um.log.Errorf("Can't gets an user (%s) from the database", userName)
		return UserRecord{}, errors.New("managers-user: can't get user from the database")
	}

	var user UserRecord
	query := `SELECT id, username, password_hash FROM users WHERE username = $1`
	if err := um.db.QueryRow(query, userName).Scan(&user.Id, &user.UserName, &user.Password); err != nil {
		um.log.Errorf("Can't gets an user (%s) from the database: %v", userName, err)
		return UserRecord{}, errors.New("managers-user: can't get user from the database")
	}

	um.log.Infof("User %s getted successfully", userName)
	return user, nil
}

func (um *userManger) GetUserById(userId string) (UserRecord, error) {
	if userId == "" {
		um.log.Errorf("Can't gets an user (%s) from the database", userId)
		return UserRecord{}, errors.New("managers-user: can't get user from the database")
	}

	var user UserRecord
	query := `SELECT id, username, password_hash FROM users WHERE id = $1`
	if err := um.db.QueryRow(query, userId).Scan(&user.Id, &user.UserName, &user.Password); err != nil {
		um.log.Errorf("Can't gets an user (%s) from the database: %v", userId, err)
		return UserRecord{}, errors.New("managers-user: can't get user from the database")
	}

	um.log.Infof("User %s getted successfully", userId)
	return user, nil
}
