package main

import "fmt"

func main() {
	num1 := 6
	num2 := 8
	fmt.Printf("交换前num1的值为：%d\n", num1)
	fmt.Printf("交换前num2的值为：%d\n", num2)
	//exchange1(num1, num2)
	exchange3(&num1, &num2)
	fmt.Printf("交换后num1的值为：%d\n", num1)
	fmt.Printf("交换后num2的值为：%d\n", num2)
}

func exchange1(x, y int) int {
	var tmp int
	tmp = x
	x = y
	y = tmp
	return tmp
}

func exchange3(x *int, y *int) int {
	var tmp int
	tmp = *x
	*x = *y
	*y = tmp
	return tmp
}
