package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 3)

	timeout := make(chan bool, 1)

	go func() {
		ch <- 6
		ch <- 7
		ch <- 8
	}()

	go func() {
		time.Sleep(6)
		timeout <- true
	}()

	select {
	case <-ch:
		fmt.Println(<-ch)
	case <-timeout:
		fmt.Println(<-timeout)
	}

}
