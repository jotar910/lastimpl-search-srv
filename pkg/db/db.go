package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBname   string
}

// GetConnection ...
func GetConnection(config *DBConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DBname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
