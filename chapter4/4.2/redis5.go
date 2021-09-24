package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "121.4.147.32:6379", redis.DialPassword("zah56641189"))
	if err != nil {
		fmt.Println("conn redis failed,err:", err)
		return
	}
	defer conn.Close()
	conn.Send("MULTI")
	conn.Send("INCR", "foo")
	conn.Send("INCR", "bar")
	r, err := conn.Do("EXEC")
	fmt.Println(r)
}
