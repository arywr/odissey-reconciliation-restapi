package app

import (
	"database/sql"
	"odissey-golang/odissey-reconciliation-restapi/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/od_go_reconciliation")
	helper.PanicIfError(err)

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(4)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
