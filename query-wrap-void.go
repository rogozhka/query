package query

import "database/sql"

type queryWrapVoid struct {
}

func NewQueryWrapVoid(db *sql.DB) *queryWrapVoid {
	return &queryWrapVoid{}
}

func (p *queryWrapVoid) FetchLimited(strQuery string, rowsLimit uint64) ([]Row, error) {
	return []Row{}, nil
}

func (p *queryWrapVoid) Fetch(strQuery string) ([]Row, error) {
	return p.FetchLimited(strQuery, 0)
}

func (p *queryWrapVoid) Select(strQuery string) ([]Row, error) {
	return p.Fetch(strQuery)
}

func (p *queryWrapVoid) SelectLimited(strQuery string, rowsLimit uint64) ([]Row, error) {
	return p.FetchLimited(strQuery, rowsLimit)
}

func (p *queryWrapVoid) Exec(strQuery string, args ...interface{}) (sql.Result, error) {
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
