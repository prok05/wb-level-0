package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prok05/wb-level-0/config"
	"log"
)

type PgStorage struct {
	pool *pgxpool.Pool
}

func NewPgStorage() *PgStorage {
	databaseUrl := fmt.Sprintf("%s://%s:%s@%s/%s",
		config.Envs.PublicHost,
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBAddress,
		config.Envs.DBName)

	dbConfig, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		log.Fatalf("error while parsing pgx config %v", err)
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatalf("error while creating pgx pool %v", err)
	}

	err = dbpool.Ping(context.Background())
	if err != nil {
		log.Fatalf("error while connecting to db: %v", err)
	}

	log.Println("DB: Successfully connected")

	return &PgStorage{pool: dbpool}
}

func (s *PgStorage) Init() (*pgxpool.Pool, error) {
	if err := s.createrOrdersTable(); err != nil {
		return nil, err
	}
	log.Println("DB: Successfully initialized")
	return s.pool, nil
}

func (s *PgStorage) createrOrdersTable() error {
	_, err := s.pool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS orders (
		    order_uid VARCHAR(255) PRIMARY KEY,
		    track_number VARCHAR(255),
		    entry VARCHAR(255),
		    delivery_info JSONB,
		    payment_info JSONB,
		    items JSONB,
		    locale VARCHAR(255),
		    internal_signature VARCHAR(255),
		    customer_id VARCHAR(255),
		    delivery_service VARCHAR(255),
		    shardkey VARCHAR(255),
		    sm_id BIGINT,
		    date_created TIMESTAMP,
		    oof_shard VARCHAR(255)
		)
	`)
	return err
}
