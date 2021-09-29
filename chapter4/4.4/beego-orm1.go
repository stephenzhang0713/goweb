package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	maxIdle := 30
	maxConn := 30
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:zah56641189@tcp(121.4.147.32:3306)/chapter4?charset=utf8", orm.MaxIdleConnections(maxIdle), orm.MaxOpenConnections(maxConn))
	orm.RegisterModel(new(BeegoUser))
	//orm.Debug = true
}

type BeegoUser struct {
	Id    int //默认主健为id
	Name  string
	Phone string
}

func main() {
	//var w io.Writer
	//orm.DebugLog = orm.NewLog(w)
	o := orm.NewOrm()
	user := new(BeegoUser)
	user.Name = "Shirdon1"
	user.Phone = "18888888889"
	fmt.Println(o.Insert(user))

	//查询+++++++
	user1 := &BeegoUser{}
	user1.Id = 1
	// 等价sql: select * from beego_user where id = 6
	err := o.Read(user1, "id")
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user1.Id, user1.Name)
	}

	user2 := &BeegoUser{}
	user2.Id = 2
	user2.Name = "James"
	num, err1 := o.Update(user2)
	if err1 != nil {
		fmt.Println("更新失败")
	} else {
		fmt.Println("更新数据影响的行数:", num)
	}

	user3 := &BeegoUser{}
	user3.Id = 2
	num1, err2 := o.Delete(user3)
	if err2 != nil {
		fmt.Println("删除失败")
	} else {
		fmt.Println("删除数据影响的行数:", num1)
	}

	to, err5 := o.Begin()
	if err5 != nil {
		logs.Error("start the transaction failed")
	}
	user4 := BeegoUser{}
	user4.Id = 3
	user4.Name = "James"

	user5 := BeegoUser{}
	user5.Id=12
	user5.Name="Wade"
	_, err3 := o.Update(&user4)
	_, err4 := o.Update(&user5)
	if err3 != nil || err4 != nil {
		to.Rollback()
	} else {
		to.Commit()
	}
}
