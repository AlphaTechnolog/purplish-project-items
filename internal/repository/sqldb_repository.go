package repository

// SQLDBRepository interface for executing arbitrary SQL queries on any SQL client.
type SQLDBRepository interface {
	Query(query string, args ...any) (Rows, error)
	QueryRow(query string, args ...any) Row
	Execute(query string, args ...any) (Result, error)
}

type Rows interface {
	Next() bool
	Scan(dest ...any) error
	Close() error
}

type Row interface {
	Scan(dest ...any) error
}

type Result interface {
	LastInsertedID() (int64, error)
	RowsAffected() (int64, error)
}
