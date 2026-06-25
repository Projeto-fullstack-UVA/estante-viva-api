package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"
	"github.com/jackc/pgx/v5"
)

func scanEvent(row pgx.Row) (*entities.Event, error) {
	var e entities.Event
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Date, &e.Location, &e.InstitutionId, &e.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &e, nil
}

func GetEvents(ctx context.Context) ([]entities.Event, error) {
	rows, err := Pool.Query(ctx,
		`SELECT id, name, description, date, location, institution_id, created_at FROM events`)
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

func GetEventById(ctx context.Context, id int64) (*entities.Event, error) {
	row := Pool.QueryRow(ctx,
		`SELECT id, name, description, date, location, institution_id, created_at FROM events
		WHERE id = $1`, id)
	return scanEvent(row)
}

func CreateEvent(ctx context.Context, event entities.Event) (int64, error) {
	result, err := Pool.Exec(ctx,
		`INSERT INTO events (name, description, date, location, institution_id, created_at) VALUES
		($1, $2, $3, $4, $5, $6)`,
		event.Name, event.Description, event.Date, event.Location, event.InstitutionId, time.Now())
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

func DeleteEvent(ctx context.Context, id int64) (int64, error) {
	result, err := Pool.Exec(ctx,
		`DELETE FROM events WHERE id = $1`, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}
