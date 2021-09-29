package model

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type BeegoUser struct {
	Id    int //默认主健为id
	Name  string
	Phone string
}

func init() {
	orm.RegisterModel(new(BeegoUser))
}