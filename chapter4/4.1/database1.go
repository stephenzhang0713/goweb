package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:zah56641189@tcp(121.4.147.32:3306)/chapter4")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
