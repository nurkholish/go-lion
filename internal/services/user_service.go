package services

import (
	"lion_parcel/internal/models"
	"lion_parcel/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo             *repositories.UserRepository
	BlacklistedTokenRepo *repositories.BlacklistedTokenRepository
}

func NewUserService(userRepo *repositories.UserRepository, blacklistedTokenRepo *repositories.BlacklistedTokenRepository) *UserService {
	return &UserService{
		UserRepo:             userRepo,
		BlacklistedTokenRepo: blacklistedTokenRepo,
	}
}

func (s *UserService) RegisterUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	if user.Role == "" {
		user.Role = "user"
	}

	return s.UserRepo.Create(user)
}

func (s *UserService) LogoutUser(token string) error {
	return s.BlacklistedTokenRepo.AddToBlacklist(token)
}

func (s *UserService) AuthenticateUser(username, password string) (*models.User, error) {
	user, err := s.UserRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	// Compare the hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}
