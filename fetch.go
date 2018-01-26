//
// query is a tiny package for executing select-like queries and converting sql.Rows into slice of map[string]string
// The new map[string]string is allocated for each row. Keys are column names as seen in sql.
//
package query

import (
	"database/sql"
	"fmt"
)

//
// Row represents single line of database result
//
// Note: when you pass a map "by value"
// actually you pass a pointer to map struct
//
type Row map[string]string

//
// Fetch is used to execute sql query
// and convert expected rows to array of map[string]string
// with keys as a column names
//
func Fetch(strQuery string, db *sql.DB) ([]Row, error) {
	return FetchLimited(strQuery, 0, db)
}

//
// FetchLimited is used to execute sql query
// and convert expected rows to array of map[string]string
// with keys as a column names; no more than rowsLimit elements
// OR unlimited if 0
//
func FetchLimited(strQuery string, rowsLimit uint64, db *sql.DB) ([]Row, error) {

	var arrRes []Row

	dbRows, err := db.Query(strQuery)
	if err != nil {
		return nil, err
	}
	if dbRows == nil {
		// cannot defer potentially nil.Close()
		return nil, fmt.Errorf("Query reutrned nil *Rows:%s", strQuery)
	}
	defer dbRows.Close()

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

		//
		// Pointer to map is appended
		// to the res. slice;
		// next getRow will allocate new map
		//
		arrRes = append(arrRes, row)
		i++
	}

	return arrRes, nil
}

//
// Select is a Fetch alias
//
func Select(strQuery string, db *sql.DB) ([]Row, error) {
	return Fetch(strQuery, db)
}
