package db

import (
	"context"
	"fmt"
	"main/apperrors"
	"main/config"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type postgresSqlDatabaseProvider struct {
	Db *sql.DB
}

func (p *postgresSqlDatabaseProvider) GetDb() *sql.DB {
	return p.Db
}

func (db *postgresSqlDatabaseProvider) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	tx, ok := ExtractTxFromContext(ctx)
	if !ok {
		return db.Db.QueryRowContext(ctx, query, args...)
	}
	return tx.QueryRowContext(ctx, query, args...)
}

func (db *postgresSqlDatabaseProvider) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	tx, ok := ExtractTxFromContext(ctx)
	if !ok {
		return db.Db.QueryContext(ctx, query, args...)
	}
	return tx.QueryContext(ctx, query, args...)
}

func (db *postgresSqlDatabaseProvider) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	tx, ok := ExtractTxFromContext(ctx)
	if !ok {
		return db.Db.ExecContext(ctx, query, args...)
	}
	return tx.ExecContext(ctx, query, args...)
}

func (db *postgresSqlDatabaseProvider) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	tx, ok := ExtractTxFromContext(ctx)
	if !ok {
		return db.Db.PrepareContext(ctx, query)
	}
	return tx.PrepareContext(ctx, query)
}

func NewPostgresSqlDatabaseProvider(cfg *config.Database) (SqlDatabaseProvider, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(172.17.0.1:3306)/%s",
		cfg.User,
		cfg.Password,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, &apperrors.ErrInternal{
			Message: fmt.Sprintf("failed to connect to mysql database: %v", err.Error())}
	}

	err = db.Ping()
	if err != nil {
		return nil, &apperrors.ErrInternal{
			Message: fmt.Sprintf("ping to postgres mysql failed: %v", err.Error())}
	}

	return &postgresSqlDatabaseProvider{Db: db}, nil
}
