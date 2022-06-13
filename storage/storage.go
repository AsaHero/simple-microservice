package storage

import (
	"github.com/AsaHero/simple-microservice/storage/postgres"
	"github.com/AsaHero/simple-microservice/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI struct {
	repo.LaptopRepository
}

func NewStorage(db *sqlx.DB) *StorageI {
	return &StorageI{
		LaptopRepository: postgres.NewLaptopPostgres(db),
	} 
}
