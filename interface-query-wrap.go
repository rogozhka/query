package query

import "database/sql"

// queryWrapInterface is a summary.
type queryWrapInterface interface {
	FetchLimited(strQuery string, rowsLimit uint64) ([]map[string]string, error)
	Fetch(strQuery string) ([]map[string]string, error)
	Select(strQuery string) ([]map[string]string, error)
	SelectLimited(strQuery string, rowsLimit uint64) ([]map[string]string, error)
	Exec(strQuery string, args ...interface{}) (sql.Result, error)
}
