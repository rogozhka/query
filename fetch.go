package query

import (
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

// FetchLimited is used to execute sql query
// and convert expected rows to array of map[string]string
// with keys as a column names; no more than rowsLimit elements
// OR unlimited if 0.
func FetchLimited(strQuery string, rowsLimit uint64, db *sql.DB) ([]map[string]string, error) {

	rows, err := db.Query(strQuery)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, fmt.Errorf("%w | %s", ErrResult, strQuery)
	}
	defer rows.Close()

	return ScanLimited(rows, rowsLimit)
}

// Select is a Fetch alias.
func Select(strQuery string, db *sql.DB) ([]map[string]string, error) {
	return Fetch(strQuery, db)
}
