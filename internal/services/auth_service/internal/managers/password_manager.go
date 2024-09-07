package managers

import (
	"errors"
	"strconv"

	"github.com/peygy/medapp/internal/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

type IPasswordManager interface {
	HashPassword(password string) (string, error)
	ValidPassword(password string) error
	CheckPasswordHash(password, hash string) error
}

type passwordManager struct {
	minLen int
	log    logger.ILogger
}

func NewPasswordManager(log logger.ILogger) IPasswordManager {
	log.Info("PasswordManager created")
	return &passwordManager{minLen: 7, log: log}
}

func (p passwordManager) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		p.log.Errorf("Can't create hashed password with error: %v", err)
		return "", errors.New("managers-password: can't create hashed password")
	}

	p.log.Info("Password is hashed successfully")
	return string(bytes), nil
}

func (p passwordManager) ValidPassword(password string) error {
	if len(password) < p.minLen {
		p.log.Error("Password length less than minimum length")
		return errors.New("managers-password: user password is not valid: password length less than " + strconv.Itoa(p.minLen))
	}

	p.log.Info("Password is valided")
	return nil
}

func (p passwordManager) CheckPasswordHash(password, hash string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		p.log.Errorf("Password is not correct: %v", err)
		return errors.New("managers-password: password is not correct")
	}

	return nil
}
