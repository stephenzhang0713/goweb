package main

import "github.com/gin-gonic/gin"

type User1 struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
}

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		u := &User1{
			Name:  "asv",
			Email: "aada@qq.com",
		}
		c.XML(200, u)
	})
	r.Run()
}
