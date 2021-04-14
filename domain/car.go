package domain

import (
	"fmt"
)

type Car struct {
	Base
	Model string
	UserId uint
}

type CarRepository interface {
	GetById(id uint) (Car, error)
	GetByUserId(userId uint) ([]Car, error)
	Create(car Car) (*Car, error)
}

func (c Car) String() string {
	return fmt.Sprintf("Car<#%d model: %s user: %d>", c.Id, c.Model, c.UserId)
}