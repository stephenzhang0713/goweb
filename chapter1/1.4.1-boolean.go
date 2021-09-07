package main

import "fmt"

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() {
	fmt.Println(boolToInt(true))
}
