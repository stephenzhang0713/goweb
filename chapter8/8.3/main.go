package main

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

var (
	db            *gorm.DB
	sqlConnection = "root:zah56641189@tcp(127.0.0.1:3306)/chapter8?" +
		"charset=utf8&parseTime=true"
)

type (
	// 数据表结构体类
	User struct {
		ID       uint   `json:"id"`
		Phone    string `json:"phone"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	// 响应返回的结构体
	UserRes struct {
		ID    uint   `json:"id"`
		Phone string `json:"phone"`
		Name  string `json:"name"`
	}
)

// 初始化
func init() {
	// 打开数据库连接
	var err error
	db, err = gorm.Open(mysql.Open(sqlConnection), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
}

// md5加密
func md5Password(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	router := gin.Default()
	v2 := router.Group("/api/v2/user")
	{
		v2.POST("/", createUser)
		v2.GET("/", fetchAllUser)
		v2.GET("/:id", fetchUser)
		v2.PUT("/:id", updateUser)
		v2.DELETE("/:id", deleteUser)
	}
	router.Run()
}

// 创建新用户
func createUser(c *gin.Context) {
	phone := c.PostForm("phone")
	name := c.PostForm("name")

	user := User{
		Phone:    phone,
		Name:     name,
		Password: md5Password("66666666"),
	}
	db.Save(&user)
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "User created successfully",
		"ID":      user.ID,
	})

}

// 获取所有用户
func fetchAllUser(c *gin.Context) {
	var user []User
	var _userRes []UserRes

	db.Find(&user)

	if len(user) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No user found!",
		})
		return
	}

	for _, item := range user {
		_userRes = append(_userRes, UserRes{
			ID:    item.ID,
			Phone: item.Phone,
			Name:  item.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   _userRes,
	})
}

func fetchUser(c *gin.Context) {
	var user User
	ID := c.Param("id")

	db.First(&user, ID)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No user found!",
		})
		return
	}

	res := UserRes{
		ID:    user.ID,
		Phone: user.Phone,
		Name:  user.Name,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   res,
	})
}

func updateUser(c *gin.Context) {
	var user User
	userId := c.Param("id")
	db.First(&user, userId)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No user found!",
		})
		return
	}

	db.Model(&user).Update("phone", c.PostForm("phone"))
	db.Model(&user).Update("name", c.PostForm("name"))
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Updated user successfully!",
	})
}

func deleteUser(c *gin.Context) {
	var user User
	userId := c.Param("id")

	db.First(&user, userId)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusOK,
			"message": "No user found!",
		})
		return
	}

	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User deleted successfully!",
	})
}
