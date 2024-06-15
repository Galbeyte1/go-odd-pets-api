package db

import (
	"errors"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

var (
	ErrConnectingDatabase = errors.New("failed to connect to database")
	ErrNotImplemented     = errors.New("not implemented")
)

type Database struct {
	Client *sqlx.DB
}

func NewDatabase() (*Database, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)

	dbConnection, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return &Database{}, fmt.Errorf("could not connect to the database: %s", err)
	}

	return &Database{
		Client: dbConnection,
	}, nil
}
