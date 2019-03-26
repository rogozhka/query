package query

import (
	"database/sql"
	"fmt"
)

//
// columns object represents query result headers
// and provides place to extract sql.Row data
//
type columns struct {

	//
	// column data receiver
	// as array of *sql.RawBytes
	//
	columnPointers []interface{}

	//
	// Number of columns in sql result
	//
	colCount int

	//
	// Column names used as row's keys
	//
	colNames []string
}

//
// newScanColumns returns initialized columns entity
// for one single query
//
func newScanColumns(colNames []string) *columns {
	namesNum := len(colNames)

	res := &columns{
		columnPointers: make([]interface{}, namesNum),
		colCount:       namesNum,
		colNames:       colNames,
	}
	for i := 0; i < res.colCount; i++ {
		res.columnPointers[i] = new(sql.RawBytes)
	}

	return res
}

//
// getRow extracts map[string]string known as Row
// from query result known as sql.Rows
//
func (p *columns) getRow(rows *sql.Rows) (Row, error) {

	//
	// Every new row should allocate it's own map
	//
	row := make(map[string]string, p.colCount)

	if err := rows.Scan(p.columnPointers...); err != nil {
		return nil, err
	}

	for i := 0; i < p.colCount; i++ {

		if rb, ok := p.columnPointers[i].(*sql.RawBytes); ok {

			row[p.colNames[i]] = string(*rb)

		} else {
			return nil, fmt.Errorf("cannot convert index %d column %v to type *sql.RawBytes", i, p.colNames[i])
		}
	}
	return row, nil
}
