package database

import (
	"database/sql"
	"fmt"

	"github.com/alphatechnolog/purplish-items/internal/repository"
)

type SQLiteRepository struct {
	db *sql.DB
}

type sqlRows struct {
	*sql.Rows
}

func (r *sqlRows) Close() error {
	return r.Rows.Close()
}

type sqlRow struct {
	*sql.Row
}

type sqlResult struct {
	sql.Result
}

func (r *sqlResult) LastInsertedID() (int64, error) {
	return r.Result.LastInsertId()
}

func (r *sqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

func NewSQLiteRepository(db *sql.DB) repository.SQLDBRepository {
	return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) Query(query string, args ...any) (repository.Rows, error) {
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("sqliteRepository query error: %w", err)
	}
	return &sqlRows{Rows: rows}, nil
}

func (r *SQLiteRepository) QueryRow(query string, args ...any) repository.Row {
	row := r.db.QueryRow(query, args...)
	return &sqlRow{Row: row}
}

func (r *SQLiteRepository) Execute(query string, args ...any) (repository.Result, error) {
	result, err := r.db.Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("sqliteRepository execute error: %w", err)
	}

	return &sqlResult{Result: result}, nil
}
