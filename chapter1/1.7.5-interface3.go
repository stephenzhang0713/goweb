package main

import "fmt"

func main() {
	var a interface{} = func(a int) string {
		return fmt.Sprintf("d:%d", a)
	}
	switch b := a.(type) {
	case nil:
		fmt.Println("nil")
	case *int:
		fmt.Println(*b)
	case func(int) string:
		fmt.Println(b(66))
	case fmt.Stringer:
		fmt.Println(b)
	default:
		fmt.Println("unknown")
	}
}
