
package repositories

import (
	"context"
	"database/sql"
	"errors"

	"devora/internal/models"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		DB: db,
	}
}

func (r *AuthRepository) CheckUser(
	ctx context.Context,
	email string,
) (models.User, error) {
	query := `
		SELECT id, username, email, password
		FROM users
		WHERE email = ?
	`

	var user models.User

	err := r.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, sql.ErrNoRows
		}

		return models.User{}, err
	}

	return user, nil
}
