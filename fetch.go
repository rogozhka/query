package query

import (
	"context"
	"database/sql"
	"fmt"
)

// Fetch is used to execute sql
// and convert expected rows
// to slice of map[string]string
// with keys as a column names.
func Fetch(strQuery string, db *sql.DB) ([]map[string]string, error) {
	return FetchLimited(strQuery, 0, db)
}

// Fetch is used to execute sql
// and convert expected rows
// to slice of map[string]string
// with keys as a column names.
func FetchContext(ctx context.Context, strQuery string, db *sql.DB) ([]map[string]string, error) {
	return FetchLimitedContext(ctx, strQuery, 0, db)
}

// FetchLimited is used to execute sql query
// and convert expected rows to array of map[string]string
// with keys as a column names; no more than rowsLimit elements
// OR unlimited if 0.
func FetchLimited(strQuery string, rowsLimit uint64, db *sql.DB) ([]map[string]string, error) {
	return FetchLimitedContext(context.Background(), strQuery, rowsLimit, db)
}

// Select is a Fetch alias.
func Select(strQuery string, db *sql.DB) ([]map[string]string, error) {
	return Fetch(strQuery, db)
}

// FetchLimited is used to execute sql query
// and convert expected rows to array of map[string]string
// with keys as a column names; no more than rowsLimit elements
// OR unlimited if 0.
func FetchLimitedContext(ctx context.Context, query string, maxRows uint64, db *sql.DB) ([]map[string]string, error) {
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, fmt.Errorf("%w | %s", ErrResult, query)
	}
	defer rows.Close()
	return ScanLimitedContext(ctx, rows, maxRows)
}
