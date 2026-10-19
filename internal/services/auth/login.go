
package auth_services

import (
	"context"
	"database/sql"
	"errors"

	"devora/internal/utils"

	"golang.org/x/crypto/bcrypt"

	"strconv"
)

type LoginResponse struct {
	ID           int    `json:"id"`
	Username     string `json:"name"`
	Email        string `json:"email"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (s *AuthService) VerifyLogin(
	ctx context.Context,
	email string,
	password string,
) (LoginResponse, error) {
	foundUser, err := s.AuthRepo.CheckUser(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return LoginResponse{}, ErrInvalidCredentials
		}

		return LoginResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(foundUser.Password),
		[]byte(password),
	); err != nil {
		return LoginResponse{}, ErrInvalidCredentials
	}

	userID := strconv.FormatInt(foundUser.ID, 10)

	accessToken, err := utils.GenerateAccessToken(userID)
	if err != nil {
		return LoginResponse{}, err
	}

	refreshToken, err := utils.GenerateRefreshToken(userID)
	if err != nil {
		return LoginResponse{}, err
	}

	var response LoginResponse

	response.ID = int(foundUser.ID)
	response.Username = foundUser.Username
	response.Email = foundUser.Email
	response.AccessToken = accessToken
	response.RefreshToken = refreshToken

	return response, nil
}
