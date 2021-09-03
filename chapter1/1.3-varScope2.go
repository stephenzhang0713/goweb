package main

import "fmt"

//声明全局变量
var global int

func main() {
	//声明局部变量
	var local1, local2 int

	//初始化参数
	local1 = 8
	local2 = 10
	global = local1 * local2

	fmt.Printf("local1 = %d, local2 = %d and g = %d\n", local1, local2, global)
}
