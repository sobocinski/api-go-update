package pgsql

import (
	"github.com/sobocinski/api-go-update/domain"
	"github.com/go-pg/pg/v10"
)

type DbCarRepository struct {
	db *pg.DB
}

// NewUserRepository - returns DB(sql) UserRepository
func NewCarRepository(db *pg.DB) *DbCarRepository {
	return &DbCarRepository{
		db: db,
	}
}

func (r *DbCarRepository) GetById(id uint) (domain.Car, error) {
	user := domain.Car{}
	err := r.db.Model(&user).Where("? = ?", pg.Ident("id"), id).Select()
	return user, err
}

func (r *DbCarRepository) GetByUserId(userId uint) ([]domain.Car, error) {
	var cars []domain.Car
	err := r.db.Model(&cars).Where("user_id = ?", userId).Select()
	return cars, err
}

func (r *DbCarRepository) Create(car domain.Car) (*domain.Car, error) {
	_, err := r.db.Model(&car).Returning("id").Returning("created_at").Insert()
	return &car, err
}
