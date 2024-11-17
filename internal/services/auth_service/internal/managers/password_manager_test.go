package managers

import (
	"testing"

	"github.com/peygy/medapp/internal/pkg/logger"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestNewPasswordManager(t *testing.T) {
	log, _ := logger.NewLogger()
	pm := NewPasswordManager(log)
	assert.NotNil(t, pm, "PasswordManager не должен быть nil")
}

func TestHashPassword(t *testing.T) {
	log, _ := logger.NewLogger()

	pm := NewPasswordManager(log)
	password := "securePassword"

	hash, err := pm.HashPassword(password)
	assert.NoError(t, err, "managers-password: can't create hashed password")
	assert.NotEmpty(t, hash, "managers-password: can't create hashed password")
}

func TestValidPassword(t *testing.T) {
	log, _ := logger.NewLogger()

	pm := NewPasswordManager(log)
	shortPassword := "123"
	validPassword := "securePassword"

	err := pm.ValidPassword(shortPassword)
	assert.Error(t, err, "managers-password: can't create hashed password")

	err = pm.ValidPassword(validPassword)
	assert.NoError(t, err, "Password is valided")
}

func TestCheckPasswordHash(t *testing.T) {
	log, _ := logger.NewLogger()

	pm := NewPasswordManager(log)
	password := "securePassword"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	err := pm.CheckPasswordHash(password, string(hash))
	assert.NoError(t, err, "Корректный пароль должен пройти проверку")

	wrongPassword := "wrongPassword"
	err = pm.CheckPasswordHash(wrongPassword, string(hash))
	assert.Error(t, err, "Некорректный пароль должен возвращать ошибку")
}
