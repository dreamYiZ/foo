package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func initDB() (db *sql.DB){
	db, err := sql.Open("mysql", "ren_too_001:fdslk238HfdsUWE@/ren_too_001")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	fmt.Println("db init!")
	return db
}
