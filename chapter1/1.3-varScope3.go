package main

import "fmt"

//声明全局变量
var global1 int = 8

func main() {
	//	声明局部变量
	var global1 int = 999
	fmt.Printf("global1 = %d\n", global1)
}
