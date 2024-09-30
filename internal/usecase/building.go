package usecase

import (
	"buildings/internal/domain/model"
	"buildings/internal/domain/repository"
	"context"

	"github.com/sirupsen/logrus"
)

type Building struct {
	bldRepo repository.Building
	log     *logrus.Logger
}

func NewBuilding(bldRepo repository.Building, log *logrus.Logger) *Building {
	return &Building{
		bldRepo: bldRepo,
		log:     log,
	}
}

func (b *Building) Get(ctx context.Context, filter model.Filter) ([]*model.Building, error) {
	log := b.log.WithField("op", "internal/usecase/building/Get")

	buildings, err := b.bldRepo.GetByFilter(ctx, filter)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return buildings, nil
}

func (b *Building) Save(ctx context.Context, building model.AddBuilding) (uint64, error) {
	log := b.log.WithField("op", "internal/usecase/building/Save")

	bld := model.Building{
		Name:   building.Name,
		City:   building.City,
		Year:   building.Year,
		Floors: building.Floors,
	}

	id, err := b.bldRepo.Save(ctx, bld)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return id, nil
}
