package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db            *gorm.DB
	sqlConnection = "root:zah56641189@tcp(121.4.147.32:3306)/chapter4?" +
		"charset=utf8&parseTime=true"
)

type GormUser struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func init() {
	var err error
	db, err = gorm.Open(mysql.Open(sqlConnection), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&GormUser{})
}

//md5加密
func md5Password(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	//创建用户
	//gormUser := GormUser{
	//	Phone:    "13888888888",
	//	Name:     "Shirdon",
	//	Password: md5Password("666666"), //用户密码
	//}
	//db.Save(&gormUser)

	//var user = new(GormUser)
	//db.Where("phone = ?","13888888888").Delete(&user)

	//查询用户
	//var user = new(GormUser)
	//db.Where("phone = ?", "13888888888").Find(&user)
	////db.First(&user,"phone = ?","13888888888")
	//fmt.Println(user)

	//var user = new(GormUser)
	//err := db.Model(&user).Where("phone = ?", "13888888888").Update("phone = ?", "13888888888").Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	tx := db.Begin()
	gormUser := GormUser{
		Phone:    "13888888889",
		Name:     "Shirdon",
		Password: md5Password("666666"), //用户密码
	}
	if err := tx.Create(&gormUser); err != nil {
		tx.Rollback()
		fmt.Println(err)
	}
	db.First(&gormUser, "phone = ?", "13888888889")
	tx.Commit()
}
