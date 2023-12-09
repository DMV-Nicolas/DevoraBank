package main

import (
	"database/sql"
	"log"

	"github.com/DMV-Nicolas/DevoraBank/api"
	db "github.com/DMV-Nicolas/DevoraBank/db/sqlc"
	"github.com/DMV-Nicolas/DevoraBank/util"
	_ "github.com/golang/mock/mockgen/model"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatalf("Cannot load config: %v", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalf("Cannot create server: %v", err)
	}

	if err = server.Start(config.ServerAddress); err != nil {
		log.Fatalf("Cannot start server: %v", err)
	}
}
