package postgres_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AsaHero/simple-microservice/config"
	"github.com/AsaHero/simple-microservice/pb"
	"github.com/AsaHero/simple-microservice/sample"
	"github.com/AsaHero/simple-microservice/storage/postgres"
)

func TestCreateLaptop(t *testing.T) {
	laptop := sample.NewLaptop()
	req := &pb.CreateLaptopRequest{Laptop: laptop}
	
	db, err := postgres.NewPostgresDB(config.Config)
	assert.NoError(t, err)
	
	laptopRepo := postgres.NewLaptopPostgres(db)

	_, err = laptopRepo.CreateLaptop(req)
	assert.NoError(t, err)
}
