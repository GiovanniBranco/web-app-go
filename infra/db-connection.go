package infra

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	connStr := "postgresql://postgres:Admin@123@localhost:5432/DB_COURSES?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}

	return db
}
