package main

import "fmt"

func main() {
	fmt.Println(getNumber())
}

func getNumber() int {
	var i int
	go func() {
		i = 6
	}()
	return i
}