package repositories

import (
	"context"
	"errors"
	"time"

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

func GetInstitutions(ctx context.Context) ([]entities.Institution, error) {
	rows, err := Pool.Query(ctx,
		`SELECT id, name, abbreviation, city, address, created_at FROM institutions ORDER BY name`)
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

func GetInstitutionById(ctx context.Context, id int64) (*entities.Institution, error) {
	row := Pool.QueryRow(ctx,
		`SELECT id, name, abbreviation, city, address, created_at FROM institutions
		WHERE id = $1`, id)
	return scanInstitution(row)
}

func CreateInstitution(ctx context.Context, i entities.Institution) (int64, error) {
	result, err := Pool.Exec(ctx,
		`INSERT INTO institutions (name, abbreviation, city, address, created_at) VALUES
		($1, $2, $3, $4, $5)`, i.Name, i.Abbreviation, i.City, i.Address, time.Now())
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

func UpdateInstitution(ctx context.Context, id int64, i entities.Institution) (int64, error) {
	result, err := Pool.Exec(ctx,
		`UPDATE institutions SET name = COALESCE(NULLIF($1, ''), name),
		 abbreviation = COALESCE(NULLIF($2, ''), abbreviation),
		 city = COALESCE(NULLIF($3, ''), city),
		 address = COALESCE(NULLIF($4, ''), address)
		 WHERE id = $5`,
		i.Name, i.Abbreviation, i.City, i.Address, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

func DeleteInstitution(ctx context.Context, id int64) (int64, error) {
	result, err := Pool.Exec(ctx,
		`DELETE FROM institutions WHERE id = $1`, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}
