package query

import (
	"context"
	"database/sql"
)

type queryWrapVoid struct {
}

// NewQueryWrapVoid created noop wrapper for testing.
func NewQueryWrapVoid(db *sql.DB) *queryWrapVoid {
	return &queryWrapVoid{}
}
func (p *queryWrapVoid) FetchLimited(strQuery string, rowsLimit uint64) ([]map[string]string, error) {
	return []map[string]string{}, nil
}
func (p *queryWrapVoid) Fetch(strQuery string) ([]map[string]string, error) {
	return p.FetchLimited(strQuery, 0)
}
func (p *queryWrapVoid) Select(strQuery string) ([]map[string]string, error) {
	return p.Fetch(strQuery)
}
func (p *queryWrapVoid) SelectLimited(strQuery string, rowsLimit uint64) ([]map[string]string, error) {
	return p.FetchLimited(strQuery, rowsLimit)
}
func (p *queryWrapVoid) Exec(strQuery string, args ...interface{}) (sql.Result, error) {
	return &sqlResult{}, nil
}

func (p *queryWrapVoid) FetchLimitedContext(ctx context.Context,
	query string,
	rowsLimit uint64,
) ([]map[string]string, error) {
	return []map[string]string{}, nil
}
func (p *queryWrapVoid) FetchContext(ctx context.Context, query string) ([]map[string]string, error) {
	return p.FetchLimitedContext(ctx, query, 0)
}
func (p *queryWrapVoid) SelectContext(ctx context.Context, query string) ([]map[string]string, error) {
	return p.FetchLimitedContext(ctx, query, 0)
}
func (p *queryWrapVoid) SelectLimitedContext(
	ctx context.Context,
	query string,
	rowsLimit uint64,
) ([]map[string]string, error) {
	return p.FetchLimitedContext(ctx, query, rowsLimit)
}
func (p *queryWrapVoid) ExecContext(
	ctx context.Context,
	query string,
	args ...interface{},
) (sql.Result, error) {
	return &sqlResult{}, nil
}

type sqlResult struct {
	lastInsertID int64
	rowsAffected int64
}

func (p *sqlResult) LastInsertId() (int64, error) {
	return p.lastInsertID, nil
}
func (p *sqlResult) RowsAffected() (int64, error) {
	return p.rowsAffected, nil
}
