package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}
