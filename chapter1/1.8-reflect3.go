package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name = "Go web Program"
	fmt.Println("真实name的原始值为：", name)

	v1 := reflect.ValueOf(&name)
	v2 := v1.Elem()

	v2.SetString("Go web Program2")
	fmt.Println("通过反射对象进行更新后，真实name变为：", name)
}
