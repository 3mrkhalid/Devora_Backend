package auth_services

import (
	"context"

	"devora/internal/models"

	"golang.org/x/crypto/bcrypt"

	"errors"

	"database/sql"
)


func (s *AuthService) CheckEmail(
	ctx context.Context,
	email string,
) error {

	_, err := s.AuthRepo.CheckUser(ctx, email)
	
	if err == nil {
		return  ErrEmailAlreadyExists
	}

	if errors.Is(err, sql.ErrNoRows) {
        return nil
    }
	
	return err
}


func (s *AuthService) RegisterUser(
	ctx context.Context,
	username string,
	email string,
	password string,
) (models.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	newUser := models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}
	
	user, err := s.AuthRepo.CreateUser(ctx, newUser)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
