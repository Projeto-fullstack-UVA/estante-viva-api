package repositories

import (
	"context"
	"errors"
	"strconv"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/models"
	"github.com/jackc/pgx/v5"
)

func scanUser(row pgx.Row) (*models.User, error) {
	var (
		id int64
		u  models.User
	)
	err := row.Scan(
		&id, &u.Name, &u.Email, &u.Address, &u.Document,
		&u.Cellphone, &u.Role, &u.Campus, &u.Score, &u.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	u.ID = strconv.FormatInt(id, 10)
	return &u, nil
}

// The password column is deliberately never selected, so it never leaves the database.
func GetUsers() ([]models.User, error) {
	rows, err := Pool.Query(context.Background(),
		`SELECT id, name, email, address, document, cellphone, role, campus, score, created_at
		 FROM users ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		u, err := scanUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, *u)
	}
	return users, rows.Err()
}

func GetUserByID(id int64) (*models.User, error) {
	row := Pool.QueryRow(context.Background(),
		`SELECT id, name, email, address, document, cellphone, role, campus, score, created_at
		 FROM users WHERE id = $1`, id)
	return scanUser(row)
}

func GetUserByCredentials(email, password string) (*models.User, error) {
	row := Pool.QueryRow(context.Background(),
		`SELECT id, name, email, address, document, cellphone, role, campus, score, created_at
		 FROM users WHERE email = $1 AND password = $2`, email, password)
	return scanUser(row)
}

// CreateUser inserts a user and returns the number of rows affected.
func CreateUser(u models.User) (int64, error) {
	tag, err := Pool.Exec(context.Background(),
		`INSERT INTO users (name, email, password, address, document, cellphone, role, campus, score, created_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		u.Name, u.Email, u.Password, u.Address, u.Document,
		u.Cellphone, u.Role, u.Campus, u.Score, u.CreatedAt,
	)
	if err != nil {
		return 0, err
	}
	return tag.RowsAffected(), nil
}
