package postgres

import (
	"context"
	"fmt"
	"github.com/exaring/otelpgx"
	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Config struct {
	URL     string
	Timeout time.Duration
}

func NewMustPostgresPool(cfg *Config) *pgxpool.Pool {
	conn, err := NewPostgresPool(cfg)
	if err != nil {
		panic(err)
	}
	return conn
}
func NewPostgresPool(cfg *Config) (*pgxpool.Pool, error) {
	connConfig, err := pgxpool.ParseConfig(cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("parse postgres url config error: %w", err)
	}
	connConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxdecimal.Register(conn.TypeMap())
		return nil
	}
	connConfig.ConnConfig.Tracer = otelpgx.NewTracer()

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()
	db, err := pgxpool.NewWithConfig(ctx, connConfig)
	if err != nil {
		return nil, fmt.Errorf("connect to postgres server error: %w", err)
	}
	return db, nil
}

func NewPostgres(cfg *Config) (*pgx.Conn, error) {
	connConfig, err := pgx.ParseConfig(cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("parse postgres url config error: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()
	db, err := pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		return nil, fmt.Errorf("connect to postgres server error: %w", err)
	}
	return db, nil
}
func NewMustPostgres(cfg *Config) *pgx.Conn {
	conn, err := NewPostgres(cfg)
	if err != nil {
		panic(err)
	}
	return conn
}
