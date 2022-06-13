package postgres

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/AsaHero/simple-microservice/pb"
)

type LaptopPostgres struct {
	db *sqlx.DB
}

func NewLaptopPostgres(db *sqlx.DB) *LaptopPostgres {
	return &LaptopPostgres{
		db: db,
	}
}

func (s *LaptopPostgres) CreateLaptop(req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	resp := &pb.CreateLaptopResponse{}
	tx := s.db.MustBegin()

	queryCpu := fmt.Sprintf(`INSERT INTO %s(brand, name, number_cores, number_threads, min_ghz, max_ghz) 
								VALUES($1, $2, $3, $4, $5, $6) RETURNING id`, CPU_TABLE)

	queryScreen := fmt.Sprintf(`INSERT INTO %s(width, height, size_inch, panel, multitouch) 
								VALUES($1, $2, $3, $4, $5) RETURNING id`, SCREEN_TABLE)

	queryLaptop := fmt.Sprintf(`INSERT INTO %s(brand, name, cpu_id, ram, gpu, storage, 
								screen_id, keyboard, weight, price_usd, release_year)
								VALUES($1, $2,
								(SELECT id FROM %s c WHERE c.id = $3), $4, $5, $6,
								(SELECT id FROM %s sn WHERE sn.id = $7), $8, $9, $10, $11) RETURNING laptop_id`,
		LAPTOP_TABLE, CPU_TABLE, SCREEN_TABLE)

	var cpu_id string
	err := tx.QueryRow(queryCpu, req.Laptop.Cpu.Brand, req.Laptop.Cpu.Name, req.Laptop.Cpu.NumberCores,
		req.Laptop.Cpu.NumberThreads, req.Laptop.Cpu.MinGhz, req.Laptop.Cpu.MaxGhz).Scan(&cpu_id)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error on CreateLaptop() -> inserting cpu: %w", err)
	}
	var screen_id string
	err = tx.QueryRow(queryScreen, req.Laptop.Screen.Resolution.Width, req.Laptop.Screen.Resolution.Height,
		req.Laptop.Screen.SizeInch, req.Laptop.Screen.Panel.String(), req.Laptop.Screen.Multitouch).Scan(&screen_id)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error on CreateLaptop() -> inserting screen: %w", err)
	}

	var laptop_guid string
	memoryJson, err := json.Marshal(req.Laptop.Ram)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error on CreateLaptop() -> cannot marshal ram data: %w", err)
	}

	gpuJson, err := json.Marshal(req.Laptop.Gpus)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error on CreateLaptop() -> cannot marshal gpu data: %w", err)
	}

	storageJson, err := json.Marshal(req.Laptop.Storages)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error on CreateLaptop() -> cannot marshal storage data: %w", err)
	}

	keyboardJson, err := json.Marshal(req.Laptop.Keyboeard)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error on CreateLaptop() -> cannot marshal keyboard data: %w", err)
	}

	weightJson, err := json.Marshal(req.Laptop.Weight)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error on CreateLaptop() -> cannot marshal weight data: %w", err)
	}

	err = tx.QueryRow(queryLaptop, req.Laptop.Brand, req.Laptop.Name, cpu_id, string(memoryJson), string(gpuJson),
		string(storageJson), screen_id, string(keyboardJson), string(weightJson), req.Laptop.PriceUsd,
		req.Laptop.ReleaseYear).Scan(&laptop_guid)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error on CreateLaptop() -> inserting laptop: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("error on CreateLaptop() -> cannot commit transaction: %w", err)
	}
	resp.Uuid = laptop_guid
	return resp, nil
}
