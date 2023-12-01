package main

import (
	"database/sql"

	"github.com/DMV-Nicolas/bankapp/api"
	db "github.com/DMV-Nicolas/bankapp/db/sqlc"
	"github.com/DMV-Nicolas/bankapp/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		panic("Cannot load config: " + err.Error())
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		panic("Cannot connect to DB: " + err.Error())
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		panic("Cannot start server: " + err.Error())
	}
}
