// Package query is a tiny package for executing select-like queries and converting sql.Rows into slice of map[string]string
// The new map[string]string is allocated for each row. Keys are column names as seen in sql.
package query

import (
	"context"
	"database/sql"
)

// Scan converts already executed query result
// into a slice of Row with keys as a column names.
func Scan(rows *sql.Rows) ([]map[string]string, error) {
	return ScanLimited(rows, 0)
}

func ScanContext(ctx context.Context, rows *sql.Rows) ([]map[string]string, error) {
	return ScanLimitedContext(ctx, rows, 0)
}

// ScanLimited converts already executed query result
// into a slice of Row with keys as a column names;
// no more than rowsLimit elements
// OR unlimited if 0.
func ScanLimited(rows *sql.Rows, rowsLimit uint64) ([]map[string]string, error) {
	return ScanLimitedContext(context.Background(), rows, rowsLimit)
}

// ScanLimitedContext converts already executed query result
// into a slice of Row with keys as a column names;
// no more than rowsLimit elements
// OR unlimited if 0.
func ScanLimitedContext(ctx context.Context, rows *sql.Rows, rowsLimit uint64) ([]map[string]string, error) {
	columnNames, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	sc := newScanColumns(columnNames)

	//nolint:prealloc // there is no clue about items amount
	var res []map[string]string

	if rowsLimit > 0 {
		res = make([]map[string]string, 0, rowsLimit)
	}
	i := uint64(0)
	for rows.Next() {
		if rowsLimit > 0 && i >= rowsLimit {
			break
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		row, err := sc.getRow(rows)
		if err != nil {
			return nil, err
		}
		// Pointer to map is appended
		// to the res. slice;
		// next getRow will allocate new map
		res = append(res, row)
		i++
	}
	return res, nil
}
