package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/DMV-Nicolas/DevoraBank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatalf("Cannot load config: %v", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
