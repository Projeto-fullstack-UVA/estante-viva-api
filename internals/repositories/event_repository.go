package repositories

import (
	"context"
	"errors"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"
	"github.com/jackc/pgx/v5"
)

func scanEvent(row pgx.Row) (*entities.Event, error) {
	var e entities.Event
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.InstitutionId, &e.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &e, nil
}

func GetEvents() ([]entities.Event, error) {
	rows, err := Pool.Query(context.Background(),
		`SELECT id, name, description, location, institution_id, created_at FROM events`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []entities.Event{}

	for rows.Next() {
		e, err := scanEvent(rows)
		if err != nil {
			return nil, err
		}
		events = append(events, *e)
	}
	return events, rows.Err()
}

func GetEventById(id int64) (*entities.Event, error) {
	row := Pool.QueryRow(context.Background(),
		`SELECT id, name, description, location, institution_id, created_at FROM events
		WHERE id = $1`, id)
	return scanEvent(row)
}
