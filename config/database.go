package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/sport-commerce")
	if err != nil {
		fmt.Println(err.Error())
		return 
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return 
	}

	return db, nil
}