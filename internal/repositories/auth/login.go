
package repositories

import (
	"context"
	"database/sql"
	"errors"

	"devora/internal/models"
)

type LoginRepository struct {
	DB *sql.DB
}

func NewLoginRepository(db *sql.DB) *LoginRepository {
	return &LoginRepository{
		DB: db,
	}
}

func (r *LoginRepository) CheckUser(
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
