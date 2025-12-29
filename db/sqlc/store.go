package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	*Queries
	db *sql.DB
}

const dbSourceStore = "postgresql://root:mysecretpassword@localhost:5432/simple_bank?sslmode=disable"

func NewStore(db *sql.DB) Store {
	testDB, err := pgxpool.New(context.Background(), dbSourceStore)
	if err != nil {
		fmt.Println("cannot connect to db:", err)
		os.Exit(1)
	}
	return Store{
		Queries: New(testDB),
		db:      db,
	}
}
