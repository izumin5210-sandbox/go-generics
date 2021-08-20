package main

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func sqlxSelect[T any](ctx context.Context, q sqlx.QueryerContext, query string, args ...interface{}) ([]T, error) {
	results := []T{}

	err := sqlx.SelectContext(ctx, q, &results, query, args...)
	if err != nil {
		return nil, err
	}

	return results, nil
}
