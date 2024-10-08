package usecase

import (
	"buildings/internal/domain/model"
	"context"
)

type Building interface {
	Get(ctx context.Context, filter model.Filter) ([]*model.Building, error)
	Save(ctx context.Context, building model.AddBuilding) (uint64, error)
}
