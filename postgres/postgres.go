package postgres

import (
	"context"
	"fmt"
	"github.com/aibotsoft/livebot/pkg/config"
	"github.com/exaring/otelpgx"
	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewMustPostgresPool(cfg *config.Config) *pgxpool.Pool {
	conn, err := NewPostgresPool(cfg)
	if err != nil {
		panic(err)
	}
	return conn
}
func NewPostgresPool(cfg *config.Config) (*pgxpool.Pool, error) {
	connConfig, err := pgxpool.ParseConfig(cfg.Postgres.URL)
	if err != nil {
		return nil, fmt.Errorf("parse postgres url config error: %w", err)
	}
	//fmt.Printf("%+v", connConfig)
	connConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxdecimal.Register(conn.TypeMap())
		return nil
	}
	connConfig.ConnConfig.Tracer = otelpgx.NewTracer()

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Postgres.Timeout)
	defer cancel()
	db, err := pgxpool.NewWithConfig(ctx, connConfig)
	if err != nil {
		return nil, fmt.Errorf("connect to postgres server error: %w", err)
	}
	return db, nil
}

func NewPostgres(cfg *config.Config) (*pgx.Conn, error) {
	connConfig, err := pgx.ParseConfig(cfg.Postgres.URL)
	if err != nil {
		return nil, fmt.Errorf("parse postgres url config error: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Postgres.Timeout)
	defer cancel()
	db, err := pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		return nil, fmt.Errorf("connect to postgres server error: %w", err)
	}
	return db, nil
}
func NewMustPostgres(cfg *config.Config) *pgx.Conn {
	conn, err := NewPostgres(cfg)
	if err != nil {
		panic(err)
	}
	return conn
}
