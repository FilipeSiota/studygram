package db

import (
	"database/sql"
	"fmt"

	"github.com/filipesiota/studygram/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	config := configs.GetDBConfig()

	strConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disabled", config.Host, config.Port, config.User, config.Password, config.Database)

	conn, err := sql.Open("postgres", strConn)

	if err != nil {
		// Avoid this behavior in production
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}