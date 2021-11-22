package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type User struct {
	Phone string `json:"phone" form:"phone"`
	Age   string `json:"age" form:"age"`
}

func main() {
	r := gin.Default()
	r.POST("/user/:id", func(c *gin.Context) {
		u := User{}
		if c.ShouldBind(&u) == nil {
			log.Println(u.Phone)
			log.Println(u.Age)
		}
		c.String(200, "success")
	})
	r.Run()
}
