// query is a tiny package for executing select-like queries and converting sql.Rows into slice of map[string]string
// The new map[string]string is allocated for each row. Keys are column names as seen in sql.
package query

import (
	"database/sql"
)

// Scan converts already executed query result
// into a slice of Row with keys as a column names.
func Scan(dbRows *sql.Rows) ([]map[string]string, error) {
	return ScanLimited(dbRows, 0)
}

// ScanLimited converts already executed query result
// into a slice of Row with keys as a column names;
// no more than rowsLimit elements
// OR unlimited if 0.
func ScanLimited(dbRows *sql.Rows, rowsLimit uint64) ([]map[string]string, error) {

	arrRes := []map[string]string{}

	columnNames, err := dbRows.Columns()
	if err != nil {
		return nil, err
	}

	sc := newScanColumns(columnNames)
	i := uint64(0)
	for dbRows.Next() {
		if rowsLimit > 0 && i >= rowsLimit {
			break
		}
		row, err := sc.getRow(dbRows)
		if err != nil {
			return nil, err
		}
		// Pointer to map is appended
		// to the res. slice;
		// next getRow will allocate new map
		arrRes = append(arrRes, row)
		i++
	}
	return arrRes, nil
}
