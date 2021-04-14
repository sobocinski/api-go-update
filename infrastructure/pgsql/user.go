package pgsql

import (
	"github.com/sobocinski/api-go-update/domain"
	"github.com/go-pg/pg/v10"
)

type DbUserRepository struct {
	db *pg.DB
}

// NewUserRepository - returns DB(sql) UserRepository
func NewUserRepository(db *pg.DB) *DbUserRepository {
	return &DbUserRepository{
		db: db,
	}
}

func (r *DbUserRepository) GetById(id uint) (domain.User, error) {
	user := domain.User{}
	err := r.db.Model(&user).Where("? = ?", pg.Ident("id"), id).Select()
	return user, err
}

func (r *DbUserRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Model(&users).Select()
	return users, err
}

func (r *DbUserRepository) Create(user domain.User) (*domain.User, error) {
	_, err := r.db.Model(&user).Returning("id").Returning("created_at").Insert()
	return &user, err
}

func (r *DbUserRepository) UserWithCars(userId uint) (*domain.User, error) {
	var user domain.User
	err := r.db.Model(&user).
		Column("user.*").
		Relation("Cars").
		First()

	return &user, err
}
