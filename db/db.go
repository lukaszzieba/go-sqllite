package db

import "database/sql"

var DB *sql.DB

func InitDb() error {
	var err error
	DB, err = sql.Open("sqlite3", "./books.db")
	if err != nil {
		return err
	}
	return DB.Ping()
}
