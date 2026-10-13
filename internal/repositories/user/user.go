package repositories

import (
    "context"
    "database/sql"
    "errors"

    "devora/internal/models"
)

type UserRepository struct {
    DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{
        DB: db,
    }
}

func (r *UserRepository) FindByID(id int) (*models.User, error) {
    query := `
        SELECT id, username, email
        FROM users
        WHERE id = ?
    `

    var user models.User

    err := r.DB.QueryRowContext(context.Background(), query, id).Scan(
        &user.ID,
        &user.Username,
        &user.Email,
    )

    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, nil
        }

        return nil, err
    }

    return &user, nil
}