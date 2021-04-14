package user

import (
	"github.com/sobocinski/api-go-update/domain"
	"github.com/sobocinski/api-go-update/internal/user/command"
	"github.com/sobocinski/api-go-update/internal/user/query"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	repo domain.UserRepository
}

// NewService - return new user service
func NewService(repo domain.UserRepository) *Service {
	return &Service{
		repo: repo,
	}
}

// GetUser
func (s *Service) GetUser(query query.GetUserQuery) (domain.User, error) {
	return s.repo.GetById(query.Id)
}

// CreateUser
func (s *Service) CreateUser(command command.CreateUserCommand) error {
	user := domain.User{
		Email: command.Email,
		Lang:  command.Lang,
	}

	newUser, e := s.repo.Create(user)
	log.Info("New user created", newUser)

	return e
}

// UserWithCars - test by orm
func (s *Service) UserWithCars(userId uint) error {
	userWithCars, e := s.repo.UserWithCars(userId)
	log.Info("UserWithCars", userWithCars.Cars)

	return e
}
