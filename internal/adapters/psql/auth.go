package psql_adapters

import "context"

type PSQLRepository struct {
	s string
}

func NewPSQLRepository() *PSQLRepository {
	return &PSQLRepository{s: "Heeey"}
}

func (r PSQLRepository) CreateUser(_ context.Context) string {
	return r.s
}
