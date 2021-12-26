package dao

import (
	"database/sql"
)

var dB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/user")
	if err != nil {
		panic(err)
	}

	dB = db
}


