package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/AsaHero/simple-microservice/config"
	"github.com/AsaHero/simple-microservice/pb"
	"github.com/AsaHero/simple-microservice/service"
	"github.com/AsaHero/simple-microservice/storage"
	"github.com/AsaHero/simple-microservice/storage/postgres"
)

func main() {
	server_host := config.Config.Service_HOST
	server_port := config.Config.Service_PORT
	address := fmt.Sprintf("%s:%s", server_host, server_port)

	fmt.Printf("start server on port: %s\n", server_port)
	
	db, err := postgres.NewPostgresDB(config.Config)
	if err != nil {
		log.Fatalf("cannot connect ro db: %s", err.Error())
	}
	storageI := storage.NewStorage(db)
	laptopService := service.NewLaptopServer(storageI)
	
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopService)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("connot start server: %s", err.Error())
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start server: %s", err.Error())
	}
}
