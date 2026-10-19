package auth_services

import (
	"context"

	"devora/internal/models"

	"golang.org/x/crypto/bcrypt"

	"errors"

	"database/sql"

	"devora/internal/utils"
	
	"strconv"
)


type RegisterRequest struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

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
) (RegisterRequest, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return RegisterRequest{}, err
	}

	newUser := models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}
	
	user, err := s.AuthRepo.CreateUser(ctx, newUser)
	if err != nil {
		return RegisterRequest{}, err
	}

	userId := strconv.FormatInt(user.ID, 10)

	accessToken, err := utils.GenerateAccessToken(userId)
	if err != nil {
		return RegisterRequest{}, err
	}
	refreshToken, err := utils.GenerateRefreshToken(userId)
	if err != nil {
		return RegisterRequest{}, err
	}

	var response RegisterRequest

	response.ID = user.ID
	response.Username = user.Username
	response.Email = user.Email
	response.AccessToken = accessToken
	response.RefreshToken = refreshToken

	return response, nil	
}
