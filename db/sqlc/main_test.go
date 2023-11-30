package db

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:83postgres19@localhost:5432/sakurabank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		panic("Cannot connect to DB: " + err.Error())
	}

	testQueries = New(testDB)

	m.Run()
}
