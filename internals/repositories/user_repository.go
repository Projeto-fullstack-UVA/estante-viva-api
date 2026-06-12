package repositories

import (
	"context"
	"errors"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"
	"github.com/jackc/pgx/v5"
)

func scanUser(row pgx.Row) (*entities.User, error) {
	var (
		id int64
		u  entities.User
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
	u.ID = id
	return &u, nil
}

func GetUsers() ([]entities.User, error) {
	rows, err := Pool.Query(context.Background(),
		`SELECT id, name, email, address, document, cellphone, role, campus, score, created_at
		 FROM users ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []entities.User{}
	for rows.Next() {
		u, err := scanUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, *u)
	}
	return users, rows.Err()
}

func GetUserByID(id int64) (*entities.User, error) {
	row := Pool.QueryRow(context.Background(),
		`SELECT id, name, email, address, document, cellphone, role, campus, score, created_at
		 FROM users WHERE id = $1`, id)
	return scanUser(row)
}

// GetUserByEmail returns the user with the given email, password hash included,
// for credential verification. Returns nil when no user matches.
func GetUserByEmail(email string) (*entities.User, error) {
	row := Pool.QueryRow(context.Background(),
		`SELECT id, name, email, password, address, document, cellphone, role, campus, score, created_at
		 FROM users WHERE email = $1`, email)

	var (
		id int64
		u  entities.User
	)
	err := row.Scan(
		&id, &u.Name, &u.Email, &u.Password, &u.Address, &u.Document,
		&u.Cellphone, &u.Role, &u.Campus, &u.Score, &u.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	u.ID = id
	return &u, nil
}

func CreateUser(user entities.User) (int64, error) {
	tag, err := Pool.Exec(context.Background(),
		`INSERT INTO users (name, email, password, address, document, cellphone, role, campus, score, created_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		user.Name, user.Email, user.Password, user.Address, user.Document,
		user.Cellphone, user.Role, user.Campus, user.Score, user.CreatedAt,
	)
	if err != nil {
		return 0, err
	}
	return tag.RowsAffected(), nil
}

func UpdateUserPassword(id int64, password string) error {
	_, err := Pool.Exec(context.Background(),
		`UPDATE users SET password = $1 WHERE id = $2`,
		password, id,
	)
	return err
}