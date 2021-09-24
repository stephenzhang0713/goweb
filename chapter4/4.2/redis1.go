package main

import "C"
import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "121.4.147.32:6379", redis.DialPassword("zah56641189"))

	if err != nil {
		fmt.Println("conn redis failed, err1:", err)
		return
	}
	defer c.Close()

	_, err = c.Do("Set", "username", "jack")
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err1 := redis.String(c.Do("Get", "username"))
	if err != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(res)

	_, err2 := c.Do("MSet", "username", "james", "phone", "1888888889")
	if err2 != nil {
		fmt.Println("MSet error:", err)
		return
	}

	res1, err3 := redis.Strings(c.Do("MGet", "username", "phone"))
	if err3 != nil {
		fmt.Println("MGet error: ", err3)
		return
	}
	fmt.Println(res1)

	_, err4 := c.Do("HSet", "names", "jim", "barry")
	if err4 != nil {
		fmt.Println("hset error: ", err4)
		return
	}
	res2, err5 := redis.String(c.Do("HGet", "names", "jim"))
	if err5 != nil {
		fmt.Println("hget error:", err5)
		return
	}
	fmt.Println(res2)

	_, err6 := c.Do("expire", "names", 10)
	if err6 != nil {
		fmt.Println("expire error:", err6)
		return
	}

	//队列
	_, err7 := c.Do("lpush", "Queue", "jim", "barry", 9)
	if err7 != nil {
		fmt.Println("lpush error: ", err7)
		return
	}
	for {
		r, err8 := redis.String(c.Do("lpop", "Queue"))
		if err8 != nil {
			fmt.Println("lpop error:", err8)
			break
		}
		fmt.Println(r)
	}
	res3, err9 := redis.Int(c.Do("llen", "Queue"))
	if err9 != nil {
		fmt.Println("llen error:", err9)
		return
	}
	fmt.Println(res3)

}
