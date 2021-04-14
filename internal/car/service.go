package car

import (
	"github.com/sobocinski/api-go-update/domain"
	"github.com/sobocinski/api-go-update/internal/car/command"
	"github.com/sobocinski/api-go-update/internal/car/query"
	query2 "github.com/sobocinski/api-go-update/internal/user/query"
)

type Service struct {
	repo domain.CarRepository
}

type CommentService interface {
	GetCar(query query.GetCarQuery) (domain.User, error)
	CreateCar(user *domain.User) error
	GetUserCars(userId uint) ([]domain.Car, error)
}

// NewService - returns a new car service
func NewService(repo domain.CarRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetCar(query query.GetCarQuery) (domain.Car, error) {
	return s.repo.GetById(query.Id)
}

func (s *Service) GetUserCars(query query2.GetUserCars) ([]domain.Car, error) {
	return s.repo.GetByUserId(query.UserId)
}

func (s *Service) CreateCar(command command.CreateCarCommand) (*domain.Car, error) {
	car := domain.Car{
		Model:  command.Model,
		UserId: command.UserId,
	}
	return s.repo.Create(car)
}


