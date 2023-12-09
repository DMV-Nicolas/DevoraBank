package main

import (
	"database/sql"

	"github.com/DMV-Nicolas/DevoraBank/api"
	db "github.com/DMV-Nicolas/DevoraBank/db/sqlc"
	"github.com/DMV-Nicolas/DevoraBank/util"
	_ "github.com/golang/mock/mockgen/model"
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
	server, err := api.NewServer(config, store)
	if err != nil {
		panic("Cannot create the server" + err.Error())
	}

	if err = server.Start(config.ServerAddress); err != nil {
		panic("Cannot start server: " + err.Error())
	}
}
