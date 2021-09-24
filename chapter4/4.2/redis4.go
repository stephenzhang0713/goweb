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

	conn.Send("HSET", "students", "name", "jim", "age", "19")
	conn.Send("HSET", "students", "score", "100")
	conn.Send("HGET", "students", "age")
	conn.Flush()

	res1, err := conn.Receive()
	fmt.Printf("Receive res1:%v \n", res1)
	res2, err := conn.Receive()
	fmt.Printf("Receive res2:%v \n", res2)
	res3, err := conn.Receive()
	fmt.Printf("Receive res3:%s \n", res3)
}
