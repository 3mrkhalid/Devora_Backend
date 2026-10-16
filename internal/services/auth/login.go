
package auth_services

import (
	"context"
	"database/sql"
	"errors"

	"devora/internal/models"
	repositories "devora/internal/repositories/auth"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	LoginRepo *repositories.AuthRepository
}

func NewAuthService(repo *repositories.AuthRepository) *AuthService {
	return &AuthService{
		LoginRepo: repo,
	}
}

func (s *AuthService) VerifyLogin(
	ctx context.Context,
	email string,
	password string,
) (models.User, error) {
	foundUser, err := s.LoginRepo.CheckUser(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, errors.New("invalid email or password")
		}

		return models.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(foundUser.Password),
		[]byte(password),
	); err != nil {
		return models.User{}, errors.New("invalid email or password")
	}

	return foundUser, nil
}
