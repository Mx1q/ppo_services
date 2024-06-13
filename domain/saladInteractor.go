package domain

import (
	"context"
	"github.com/google/uuid"
)

type ISaladInteractor interface {
	Create(ctx context.Context, salad *Salad) (uuid.UUID, error)
	GetById(ctx context.Context, id uuid.UUID) (*Salad, error)
	GetAll(ctx context.Context, filter *RecipeFilter, page int) ([]*Salad, int, error)
	GetAllByUserId(ctx context.Context, id uuid.UUID) ([]*Salad, error)
	GetAllRatedByUser(ctx context.Context, userId uuid.UUID, page int) ([]*Salad, int, error)
	Update(ctx context.Context, salad *Salad) error
	DeleteById(ctx context.Context, id uuid.UUID) error
}
