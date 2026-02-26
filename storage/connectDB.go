package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB() {

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Error connecting", err)
		return
	}
	defer pool.Close()

	var name string
	err = pool.QueryRow(context.Background(), "SELECT name FROM users WHERE id=$1", 1).Scan(&name)
	if err != nil {
		fmt.Println("query error", err)
		return
	}

	fmt.Println("User name", name)

}
