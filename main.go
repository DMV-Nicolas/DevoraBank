package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/DMV-Nicolas/DevoraBank/api"
	db "github.com/DMV-Nicolas/DevoraBank/db/sqlc"
	"github.com/DMV-Nicolas/DevoraBank/gapi"
	"github.com/DMV-Nicolas/DevoraBank/pb"
	"github.com/DMV-Nicolas/DevoraBank/util"
	_ "github.com/golang/mock/mockgen/model"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}
	store := db.NewStore(conn)
	runGrpcServer(config, store)
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDevorabankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot create listener: ", err)
	}

	log.Println("Start gRPC server at", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Cannot start gRPC server: ", err)
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server: ", err)
	}

	if err = server.Start(config.ServerAddress); err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
