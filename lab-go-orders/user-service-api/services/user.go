package services

import (
	"errors"
	"regexp"
	"user-service-api/dto"
	"user-service-api/models"
	"user-service-api/storage"

	"golang.org/x/crypto/bcrypt"
)

func isValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	if !hasUpper {
		return false
	}

	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	if !hasNumber {
		return false
	}

	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	if !hasLower {
		return false
	}

	return true
}

func CreateUser(req dto.RegisterRequest) (*models.User, error) {
	if !isValidPassword(req.Password) {
		return nil, errors.New("password must be at least 8 characters long and contain an uppercase letter and a number")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("could not hash password")
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	db := storage.GetDB()
	if err := db.Create(&user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	return &user, nil
}
