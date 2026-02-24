package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func initDB() {

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println("can't connecd db")
		os.Exit(1)
	}

	defer conn.Close(context.Background())
}
