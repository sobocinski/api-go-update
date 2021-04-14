package domain

import (
	"fmt"
)

type User struct {
	Base
	Email string `pg:"type:varchar(180)"`
	Lang string `pg:"type:varchar(2)"`
	Cars []*Car `pg:"rel:has-many"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetById(id uint) (User, error)
	Create(user User) (*User, error)
 	UserWithCars(userId uint) (*User, error)
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s>", u.Id, u.Email)
}

