package auth_services

import repositories "devora/internal/repositories/auth"

type AuthService struct {
    AuthRepo *repositories.AuthRepository
}

func NewAuthService(repo *repositories.AuthRepository) *AuthService {
    return &AuthService{
        AuthRepo: repo,
    }
}