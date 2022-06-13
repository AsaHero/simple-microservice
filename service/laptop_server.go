package service

import (
	"context"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/AsaHero/simple-microservice/pb"
	"github.com/AsaHero/simple-microservice/storage"
)

type LaptopServer struct {
	storage storage.StorageI
	pb.UnimplementedLaptopServiceServer
}

func NewLaptopServer(storage *storage.StorageI) *LaptopServer {
	return &LaptopServer{
		storage:  *storage,
	}
}

func (s *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	resp := &pb.CreateLaptopResponse{}
	laptop := req.Laptop

	if len(laptop.Id) > 0 {
		if _, err := uuid.Parse(laptop.Id); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error on CreateLaptop() --> uuid.NewRandom: %w", err)
		}
		laptop.Id = id.String()
	}

	if ctx.Err() == context.Canceled {
		log.Printf("request is canceled!")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("deadline is exceeded !")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}

	resp, err := s.storage.LaptopRepository.CreateLaptop(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error on CreateLaptop() -> LaptopRespoditory: %w", err)
	}

	return resp, nil
}
