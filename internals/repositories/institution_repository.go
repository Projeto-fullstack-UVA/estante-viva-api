package repositories

import (
	"context"
	"errors"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"
	"github.com/jackc/pgx/v5"
)

func scanInstitution(row pgx.Row) (*entities.Institution, error) {
	var i entities.Institution

	err := row.Scan(&i.ID, &i.Name, &i.Abbreviation, &i.City, &i.Address, &i.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &i, nil
}

func GetInstitutions() ([]entities.Institution, error) {
	rows, err := Pool.Query(context.Background(),
		`SELECT id, name, abbreviation, city, address, created_at FROM institutions ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	institutions := []entities.Institution{}
	for rows.Next() {
		i, err := scanInstitution(rows)
		if err != nil {
			return nil, err
		}
		institutions = append(institutions, *i)
	}
	return institutions, rows.Err()
}
