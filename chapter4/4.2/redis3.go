package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "121.4.147.32:6379", redis.DialPassword("zah56641189"))
	if err != nil {
		fmt.Println("conn redis failed,err:", err)
		return
	}
	defer c.Close()

	c.Send("SET", "username1", "jim")
	c.Send("SET", "username2", "jack")

	c.Flush()

	v, err1 := c.Receive()
	fmt.Printf("v:%v,err:%v\n", v, err1)
	v, err = c.Receive()
	fmt.Printf("v:%v,err:%v\n", v, err1)

	v, err = c.Receive()
	fmt.Printf("v:%v,err:%v\n", v, err1)

}
