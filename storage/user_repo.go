package storage

import (
	"cdcat/types"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUser(ctx context.Context, pool *pgxpool.Pool, user *types.User) error {

	query := `
	INSERT INTO users (name, email, plan)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	return pool.QueryRow(ctx, query,
		user.Name,
		user.Email,
		user.Plan).Scan(&user.ID)
}

func GetUser(ctx context.Context, pool *pgxpool.Pool, user *types.User) error {

	query := `
	SELECT *
	FROM users
	WHERE id = $1
	`
	return pool.QueryRow(ctx, query, user.ID).Scan(&user.ID, &user.Name, &user.Email, &user.Plan)
}
