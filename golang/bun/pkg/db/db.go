package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func Setup() (*bun.DB, error) {
	cfg := mysql.Config{
		User:      "root",
		Passwd:    "password",
		DBName:    "db",
		ParseTime: true,
	}
	sqlDB, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return bun.NewDB(sqlDB, mysqldialect.New()), nil
}
