package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name = "Go web Program"

	v1 := reflect.ValueOf(name)
	v2 := reflect.ValueOf(&name)
	v3 := v2.Elem()
	fmt.Println("v1可写性为:", v1.CanSet())
	fmt.Println("v1可写性为:", v2.CanSet())
	fmt.Println("v3可写性为:", v3.CanSet())
}
