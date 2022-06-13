package postgres

import (
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/AsaHero/simple-microservice/config"
)

const (
	LAPTOP_TABLE   = "laptop"
	CPU_TABLE      = "cpu"
	SCREEN_TABLE   = "screen"
)

func NewPostgresDB(config config.Configuration) (*sqlx.DB, error) {
	dbUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.PostgresHost, config.PostgresPort, config.PostgresUsername,
		config.PostgresPassword, config.LaptopServiceDB, config.PostgresSSLMode)
	fmt.Println(dbUrl)
	db, err := sqlx.Connect("pgx", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("error on db connection NewPostgrsDB: %w", err)
	}

	return db, nil
}
