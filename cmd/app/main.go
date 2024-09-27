package main

import (
	"github.com/Pasca11/gRPC-Auth/internal/repository/postgres"
	"github.com/Pasca11/gRPC-Auth/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	server := grpc.NewServer()
	db, err := postgres.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	service.RegisterServer(server, db)

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	log.Println("grpc server listening on :8081")
	if err := server.Serve(l); err != nil {
		panic(err)
	}
}
