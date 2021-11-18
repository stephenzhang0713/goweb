package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		fmt.Println("开始goroutine")
		ch <- "signal"
		fmt.Println("退出goroutine")
	}()
	fmt.Println("等待goroutine")
	<-ch
	fmt.Println("完成")
}
