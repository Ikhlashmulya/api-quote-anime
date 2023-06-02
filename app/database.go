package app

import (
	"database/sql"
	"quote-anime-api/helper"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:4kWJkVilqTYs1HjQ9DkC@tcp(containers-us-west-53.railway.app:5609)/railway")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(50 * time.Minute)

	return db
}
