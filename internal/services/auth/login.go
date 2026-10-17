
package auth_services

import (
	"context"
	"database/sql"
	"errors"

	"devora/internal/models"


	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) VerifyLogin(
	ctx context.Context,
	email string,
	password string,
) (models.User, error) {
	foundUser, err := s.AuthRepo.CheckUser(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, ErrInvalidCredentials
		}

		return models.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(foundUser.Password),
		[]byte(password),
	); err != nil {
		return models.User{}, ErrInvalidCredentials
	}

	return foundUser, nil
}
