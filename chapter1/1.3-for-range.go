package main

import "fmt"

func main() {
	for key, value := range []int{0, 1, -2, -2} {
		fmt.Printf("Key:%d value:%d\n", key, value)
	}

	var str = "hi 加油"
	for key, value := range str {
		fmt.Printf("key:%d value:0x%x\n", key, value)
	}

	m := map[string]int{
		"go":  100,
		"web": 100,
	}
	for _, value := range m {
		fmt.Println(value)
	}

	c := make(chan int)
	go func() {
		c <- 7
		c <- 8
		c <- 9
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}
