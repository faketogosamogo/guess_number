package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	)

const (
	DriverName = "mysql"
	ConnectionString = "root:1234@/guess_number"
)

type DB struct {
	*sql.DB
}

func NewDB(driver, conn string)(*DB, error){
	db, err := sql.Open(driver, conn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

