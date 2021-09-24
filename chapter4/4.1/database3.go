package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db1 *sql.DB

//定义一个全局变量
var u User

type User struct {
	Uid   int
	Name  string
	Phone string
}

//初始化数据库连接
func init() {
	//db, _ = sql.Open("mysql",
	//	"root:123456@tcp(127.0.0.1:3306)/chapter4")

	db1, _ = sql.Open("mysql", "root:zah56641189@tcp(121.4.147.32:3306)/chapter4")
}

//单行测试
func queryRow() {
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db1.QueryRow("select uid,name,phone from `user` where uid = ?", 1).Scan(&u.Uid, &u.Name, &u.Phone)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("uid:%d name:%s phone:%s\n", u.Uid, u.Name, u.Phone)
}

// 查询多条数据示例
func queryMultiRow() {
	rows, err := db1.Query("select uid,name,phone from `user` where uid>?", 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		err := rows.Scan(&u.Uid, &u.Name, &u.Phone)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("uid:%d name:%s phone:%s\n", u.Uid, u.Name, u.Phone)
	}
}

func InsertRow() {
	ret, err := db1.Exec("insert into user(name, phone) VALUES (?,?)", "王老五", 13988556677)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	uid, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", uid)
}

func updateRow() {
	ret, err := db1.Exec("update user set name = ? where uid = ?", "张三", 3)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

func deleteRow() {
	ret, err := db1.Exec("delete from user where uid = ?", 2)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

// 预处理查询示例
func prepareQuery() {
	stmt, err := db1.Prepare("select uid,name,phone from `user` where uid > ?")
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&u.Uid, &u.Name, &u.Phone)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("uid:%d name:%s phone:%s\n", u.Uid, u.Name, u.Phone)
	}
}

// 预处理插入示例
func prepareInsert() {
	stmt, err := db1.Prepare("insert into user(name, phone) VALUES (?,?)")
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("barry", 18799887766)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	_, err = stmt.Exec("jim", 18988888888)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	fmt.Println("insert success.")
}

func transaction() {
	tx, err := db1.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}

	_, err = tx.Exec("update user set name='james' where uid = ?", 1)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	_, err = tx.Exec("update user set name='james' where uid=?", 3)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	tx.Commit()
	fmt.Println("exec transaction success!")
}

// sql注入示例
func sqlInject(name string) {
	sqlStr := fmt.Sprintf("select uid,name,phone from user where name = '%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	ret, err := db1.Exec(sqlStr)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("get success, affected rows:%d\n", n)
}

func main() {
	//queryRow()
	//queryMultiRow()
	//InsertRow()
	//updateRow()
	//deleteRow()
	//prepareInsert()
	//transaction()
	sqlInject("james' union select * from user #")
	prepareQuery()
}
