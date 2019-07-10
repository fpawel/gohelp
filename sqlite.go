package gohelp

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func OpenSqliteDBx(fileName string) *sqlx.DB {
	return sqlx.NewDb(OpenSqliteDB(fileName), "sqlite3")
}

func OpenSqliteDB(fileName string) *sql.DB {
	conn, err := sql.Open("sqlite3", fileName)
	if err != nil {
		panic(err)
	}
	conn.SetMaxIdleConns(1)
	conn.SetMaxOpenConns(1)
	conn.SetConnMaxLifetime(0)
	return conn
}
