package main

import "fmt"

// 生成自增的整数
func IntegerGenerator() chan int {
	var ch = make(chan int)

	// 	开启goroutine
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	return ch
}

func main() {
	generator := IntegerGenerator()
	for i := 0; i < 100; i++ {
		fmt.Println(<-generator)
	}
}
