package repository

import (
	"buildings/internal/domain/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	repository.Building
}

func NewRepositories(pool *pgxpool.Pool) *Repositories {
	return &Repositories{
		Building: NewBuilding(pool),
	}
}
