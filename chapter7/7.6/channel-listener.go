package main

import (
	"fmt"
	"time"
)

func foo(i int) chan int {
	ch := make(chan int)
	go func() {
		ch <- i
	}()
	return ch
}
func main() {
	ch1, ch2, ch3 := foo(3), foo(6), foo(9)

	ch := make(chan int)
	timeout := time.After(1 * time.Second)

	go func() {
		for isTimeout := false; !isTimeout; {
			select {
			case v1 := <-ch1:
				fmt.Printf("received %d from ch1\n", v1)
				ch <- v1
			case v2 := <-ch2:
				fmt.Printf("received %d from ch2\n", v2)
				ch <- v2
			case v3 := <-ch3:
				fmt.Printf("received %d from ch3\n", v3)
				ch <- v3
			case <-timeout:
				isTimeout = true
			}
		}
	}()

	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
}
