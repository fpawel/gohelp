package gohelp

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func MustOpenSqliteDBx(fileName string) *sqlx.DB {
	return sqlx.NewDb(MustOpenSqliteDB(fileName), "sqlite3")
}

func MustOpenSqliteDB(fileName string) *sql.DB {
	conn, err := OpenSqliteDB(fileName)
	if err != nil {
		panic(err)
	}
	return conn
}

func OpenSqliteDBx(fileName string) (*sqlx.DB, error) {
	c, err := OpenSqliteDB(fileName)
	if err != nil {
		return nil, err
	}
	return sqlx.NewDb(c, "sqlite3"), nil
}

func OpenSqliteDB(fileName string) (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", fileName)
	if err != nil {
		return nil, err
	}
	conn.SetMaxIdleConns(1)
	conn.SetMaxOpenConns(1)
	conn.SetConnMaxLifetime(0)
	return conn, err
}
