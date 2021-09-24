package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func initDB() (err error) {
	db, err = sql.Open("mysql", "root:zah56641189@tcp(121.4.147.32:3306)/chapter4")
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db1 failed, err:%v\n", err)
		return
	}
}
