package services

import (
    "errors"
    "time"
    "backend_masjid/internal/config"
    // "backend_masjid/internal/models"
    "backend_masjid/internal/repository"
    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
    adminRepo *repository.AdminUserRepository
    dkmRepo   *repository.DkmUserRepository
}

func NewAuthService(adminRepo *repository.AdminUserRepository, dkmRepo *repository.DkmUserRepository) *AuthService {
    return &AuthService{adminRepo, dkmRepo}
}

func (s *AuthService) LoginAdmin(username, password string) (string, error) {
    user, err := s.adminRepo.FindByUsername(username)
    if err != nil {
        return "", errors.New("invalid username or password")
    }

    if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
        return "", errors.New("invalid username or password")
    }

    claims := jwt.MapClaims{
        "role": "admin",
        "exp":  time.Now().Add(time.Hour * 24).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(config.Cfg.JwtSecret))
}

func (s *AuthService) LoginDKM(username, password, device string) (string, error) {
    user, err := s.dkmRepo.FindByUsername(username)
    if err != nil {
        return "", errors.New("invalid username or password")
    }

    if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
        return "", errors.New("invalid username or password")
    }

    // Default role
    role := "dkm"
    if device == "tv" {
        role = "display"
    }

    claims := jwt.MapClaims{
        "role":      role,
        "client_id": user.ClientID,
        "exp":       time.Now().Add(time.Hour * 24).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(config.Cfg.JwtSecret))
}