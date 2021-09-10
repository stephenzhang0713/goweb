package main

import "fmt"

func min(arr []int) (m int) {
	m = arr[0]
	for _, v := range arr {
		if v < m {
			m = v
		}
	}
	return
}

func compute(x, y int) (int, int) {
	return x + y, x * y
}

func main() {
	arr := []int{1, 2, 3, 4, 3}
	fmt.Println(min(arr))
	fmt.Println(compute(6, 8))
}
