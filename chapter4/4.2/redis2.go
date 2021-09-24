package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "121.4.147.32:6379", redis.DialPassword("zah56641189"))
		},
		MaxIdle:     16,
		MaxActive:   1024,
		IdleTimeout: 300,
	}
}

func main() {
	c := pool.Get()
	defer c.Close()
	_, err := c.Do("Set", "username", "jack")
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err1 := redis.String(c.Do("Get", "username"))
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(r)
}
