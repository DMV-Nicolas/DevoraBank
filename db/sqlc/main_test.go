package db

import (
	"database/sql"
	"os"
	"testing"

	"github.com/DMV-Nicolas/sakurabank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		panic("Cannot load config: " + err.Error())
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		panic("Cannot connect to DB: " + err.Error())
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
