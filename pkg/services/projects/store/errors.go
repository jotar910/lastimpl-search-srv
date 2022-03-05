package store

import "fmt"

var (
	ErrNotResult  = func() string { return "sql: no rows in result set" }
	ErrDuplicated = func(constraint string) string {
		return fmt.Sprintf("pq: duplicate key value violates unique constraint %q", constraint)
	}
)
