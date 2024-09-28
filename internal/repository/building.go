package repository

import (
	"buildings/internal/domain/model"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Building struct {
	pool *pgxpool.Pool
}

func NewBuilding(pool *pgxpool.Pool) *Building {
	return &Building{
		pool: pool,
	}
}

func (b *Building) GetByFilter(ctx context.Context, filter model.Filter) ([]*model.Building, error) {
	query := "SELECT id, name, city, year, floors FROM buildings WHERE 1=1"
	var args []interface{}
	argID := 1

	if filter.City != "" {
		query += fmt.Sprintf(" AND city = $%d", argID)
		argID++
		args = append(args, filter.City)
	}
	if filter.Year != 0 {
		query += fmt.Sprintf(" AND year = $%d", argID)
		argID++
		args = append(args, filter.Year)
	}
	if filter.Floors != 0 {
		query += fmt.Sprintf(" AND floors = $%d", argID)
		argID++
		args = append(args, filter.Floors)
	}

	rows, err := b.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var buildings []*model.Building
	for rows.Next() {
		b := model.Building{}

		err := rows.Scan(
			&b.ID,
			&b.Name,
			&b.City,
			&b.Year,
			&b.Floors,
		)
		if err != nil {
			return nil, err
		}

		buildings = append(buildings, &b)
	}

	return buildings, nil
}

func (b *Building) Save(ctx context.Context, building model.Building) (uint64, error) {
	quiry := "insert into buildings (name, city, year, floors) values ($1, $2, $3, $4) returning id"

	row := b.pool.QueryRow(
		ctx,
		quiry,
		building.Name,
		building.City,
		building.Year,
		building.Floors,
	)

	var id uint64 = 0
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
