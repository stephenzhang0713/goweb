package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type User struct {
	uid   int
	Name  string
	Phone string
}

func init() {
	DB, _ = sql.Open("mysql", "root:zah56641189@tcp(121.4.147.32:3306)/goweb")
}

func GetUser(uid int) (u User) {
	err := DB.QueryRow("select uid, name, phone from user where uid = ?", uid).Scan(&u.uid, &u.Name, &u.Phone)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n")
		return
	}
	return u
}
