package usecase

import (
	"buildings/internal/domain/usecase"
	"buildings/internal/repository"

	"github.com/sirupsen/logrus"
)

type Usecases struct {
	usecase.Building
}

func NewUsecases(repos repository.Repositories, log *logrus.Logger) *Usecases {
	return &Usecases{
		Building: NewBuilding(repos.Building, log),
	}
}
