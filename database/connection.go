package database

import (
	"database/sql"
	"fmt"
	"github.com/estebanmpa/go-pure-rest-api/config"
	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	cfg := config.GetDatabaseConfig()
	conStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.DbName)
	db, err := sql.Open(cfg.Dialect, conStr)
	
	if err != nil {
		panic(err.Error())
	}

	return db
}
