package repository

import (
	"buildings/internal/domain/model"
	"context"
)

type Building interface {
	GetByFilter(ctx context.Context, filter model.Filter) ([]*model.Building, error)
	Save(ctx context.Context, building model.Building) (uint64, error)
}
