package services

import (
	"errors"

	"book-management/internal/models"
	"book-management/internal/repositories"
	"book-management/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo   *repositories.UserRepository
	jwtManager *utils.JWTManager
}

func NewAuthService(userRepo *repositories.UserRepository, jwtManager *utils.JWTManager) *AuthService {
	return &AuthService{
		userRepo:   userRepo,
		jwtManager: jwtManager,
	}
}

func (s *AuthService) Login(req *models.LoginRequest) (*models.LoginResponse, error) {
	// Validate input
	if err := utils.ValidateStruct(req); err != nil {
		return nil, errors.New("validation failed: " + err.Error())
	}

	// Find user by username
	user, err := s.userRepo.GetByUsername(req.Username)
	if err != nil {
		return nil, errors.New("failed to find user")
	}

	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT token
	token, expiresAt, err := s.jwtManager.GenerateToken(user)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &models.LoginResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		User: models.UserInfo{
			ID:       user.ID,
			Username: user.Username,
		},
	}, nil
}

func (s *AuthService) ValidateToken(tokenString string) (*utils.JWTClaims, error) {
	return s.jwtManager.ValidateToken(tokenString)
}

func (s *AuthService) GetUserByID(id int) (*models.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("failed to get user")
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
