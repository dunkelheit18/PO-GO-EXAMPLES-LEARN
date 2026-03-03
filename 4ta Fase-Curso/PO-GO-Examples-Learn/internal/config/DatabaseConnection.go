package config

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetInstance(cfg *DBConfig) *sql.DB {
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
		)

		var err error
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}
	})
	return db
}
