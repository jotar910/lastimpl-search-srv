package queries

import (
	_ "embed"
)

var (
	//go:embed create_tables.sql
	QueryCreateTables string
	//go:embed drop_tables.sql
	QueryDropTables string
)
