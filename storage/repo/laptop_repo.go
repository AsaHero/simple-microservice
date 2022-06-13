package repo

import "github.com/AsaHero/simple-microservice/pb"

type LaptopRepository interface {
	CreateLaptop(*pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error)
}