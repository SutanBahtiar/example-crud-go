package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	DriverName     = "postgres"
	DataSourceName = "postgresql://root:root@localhost:5432/example_crud?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(DriverName, DataSourceName)
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
