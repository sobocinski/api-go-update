package domain

import (
	"context"
	"time"
)

type Base struct {
	Id uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (model *Base) BeforeInsert(ctx context.Context) (context.Context, error) {
	model.CreatedAt = time.Now()
	return ctx, nil
}

func (model *Base) BeforeUpdate(ctx context.Context) (context.Context, error) {
	model.UpdatedAt = time.Now()
	return ctx, nil
}