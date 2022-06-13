package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/AsaHero/simple-microservice/config"
	"github.com/AsaHero/simple-microservice/pb"
	"github.com/AsaHero/simple-microservice/sample"
)

func main() {
	server_host := config.Config.Service_HOST
	server_port := config.Config.Service_PORT
	address := fmt.Sprintf("%s:%s", server_host, server_port)

	fmt.Printf("client dialing server: %s\n", address)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial to server: %s", err.Error())
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()
	laptop.Id = "ad3d1784-eb9b-4e4c-b209-0767ffe7cea3"

	createLaptopReq := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 2)
	defer cancel()

	resp, err := laptopClient.CreateLaptop(ctx, createLaptopReq)
	if err != nil {
		log.Fatalf("cannot create laptop: %s", err.Error())
	}

	fmt.Printf("created laptop id: %s", resp.Uuid)
}
