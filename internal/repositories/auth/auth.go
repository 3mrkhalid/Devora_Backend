package repositories

import (
	"context"
	"database/sql"

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
		return models.User{}, err
	}

	return user, nil
}


func (r *AuthRepository) CreateUser(
    ctx context.Context,
    user models.User,
) (models.User, error) {
    query := `
        INSERT INTO users (username, email, password)
        VALUES (?, ?, ?)
    `

    result, err := r.DB.ExecContext(
        ctx,
        query,
        user.Username,
        user.Email,
        user.Password,
    )
    if err != nil {
        return models.User{}, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return models.User{}, err
    }

    user.ID = id

    return user, nil
}