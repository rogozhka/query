package query

import (
	"context"
	"database/sql"
)

type queryWrap struct {
	db *sql.DB
}

func NewQueryWrap(db *sql.DB) *queryWrap {
	return &queryWrap{db: db}
}

func (p *queryWrap) FetchLimited(strQuery string, rowsLimit uint64) ([]map[string]string, error) {
	return FetchLimited(strQuery, rowsLimit, p.db)
}

func (p *queryWrap) Fetch(strQuery string) ([]map[string]string, error) {
	return p.FetchLimited(strQuery, 0)
}

func (p *queryWrap) Select(strQuery string) ([]map[string]string, error) {
	return p.Fetch(strQuery)
}

func (p *queryWrap) SelectLimited(strQuery string, rowsLimit uint64) ([]map[string]string, error) {
	return p.FetchLimited(strQuery, rowsLimit)
}

func (p *queryWrap) Exec(strQuery string, args ...interface{}) (sql.Result, error) {
	return p.db.Exec(strQuery, args...)
}

func (p *queryWrap) FetchLimitedContext(ctx context.Context, strQuery string, rowsLimit uint64) ([]map[string]string, error) {
	return FetchLimitedContext(ctx, strQuery, rowsLimit, p.db)
}

func (p *queryWrap) FetchContext(ctx context.Context, strQuery string) ([]map[string]string, error) {
	return p.FetchLimitedContext(ctx, strQuery, 0)
}

func (p *queryWrap) SelectContext(ctx context.Context, strQuery string) ([]map[string]string, error) {
	return p.FetchContext(ctx, strQuery)
}

func (p *queryWrap) SelectLimitedContext(ctx context.Context, strQuery string, rowsLimit uint64) ([]map[string]string, error) {
	return p.FetchLimitedContext(ctx, strQuery, rowsLimit)
}

func (p *queryWrap) ExecContext(ctx context.Context, strQuery string, args ...interface{}) (sql.Result, error) {
	if nil == ctx {
		ctx = context.Background()
	}
	return p.db.ExecContext(ctx, strQuery, args...)
}
