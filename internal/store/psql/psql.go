package psql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/veliancreate/books-api/internal/config"
)

// ensure only one connection pool
var dbpool *pgxpool.Pool

type PSQLStore struct {
	conn string
}

func NewPSQLStore(config config.DBConfig) (*PSQLStore, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DbName,
	)

	return &PSQLStore{
		conn: conn,
	}, nil
}

func (store *PSQLStore) Init() error {
	if dbpool != nil {
		return nil
	}

	_, err := pgxpool.Connect(context.Background(), store.conn)
	if err != nil {
		return fmt.Errorf("could not connect to db: %w", err)
	}

	return nil
}
