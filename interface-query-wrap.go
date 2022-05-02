package query

import (
	"context"
	"database/sql"
)

// queryWrapInterface is a summary.
type queryWrapInterface interface {
	FetchLimited(query string, rowsLimit uint64) ([]map[string]string, error)
	Fetch(query string) ([]map[string]string, error)
	Select(query string) ([]map[string]string, error)
	SelectLimited(query string, rowsLimit uint64) ([]map[string]string, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// queryWrapContextInterface is a summary.
type queryWrapContextInterface interface {
	FetchLimitedContext(ctx context.Context, query string, rowsLimit uint64) ([]map[string]string, error)
	FetchContext(ctx context.Context, query string) ([]map[string]string, error)
	SelectContext(ctx context.Context, query string) ([]map[string]string, error)
	SelectLimitedContext(ctx context.Context, query string, rowsLimit uint64) ([]map[string]string, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}
