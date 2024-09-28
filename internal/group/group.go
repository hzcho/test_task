package group

import "buildings/internal/usecase"

type Groups struct {
	Building
}

func NewGroups(usecases *usecase.Usecases) *Groups {
	return &Groups{
		Building: *NewBuilding(usecases.Building),
	}
}
