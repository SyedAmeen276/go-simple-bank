package db

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDB *pgxpool.Pool

const dbSource = "postgresql://root:mysecretpassword@localhost:5432/simple_bank?sslmode=disable"

func TestMain(m *testing.M) {
	var err error

	testDB, err = pgxpool.New(context.Background(), dbSource)
	if err != nil {
		fmt.Println("cannot connect to db:", err)
		os.Exit(1)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
